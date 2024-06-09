package main

import (
	"database/sql"
	"net"
	"os"
	"sync"

	"github.com/fredouric/cheese-finder-grpc/api/cheese"
	"github.com/fredouric/cheese-finder-grpc/dataset"
	"github.com/fredouric/cheese-finder-grpc/db"
	"github.com/fredouric/cheese-finder-grpc/pb/cheesev1"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	_ "modernc.org/sqlite"
)

var (
	port   string
	dbPath string
)

var app = &cli.App{
	Name:    "cheese-finder-grpc",
	Usage:   "",
	Suggest: true,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "port",
			Value:       ":3000",
			Destination: &port,
			EnvVars:     []string{"PORT"},
		},
		&cli.StringFlag{
			Name:        "DB Path",
			Value:       "./cheese.sqlite3",
			Destination: &dbPath,
			EnvVars:     []string{"DB_PATH"},
		},
	},
	Action: func(ctx *cli.Context) error {
		log.Level(zerolog.DebugLevel)

		sqlite, err := sql.Open("sqlite", dbPath)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to open db")
		}
		queries := db.Migrate(sqlite, dbPath)

		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()

			cheeses, err := dataset.Fetch()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to fetch dataset")
			}
			log.Info().Msg("entire dataset fetched")

			if err := queries.DeleteAllCheeses(ctx.Context); err != nil {
				log.Fatal().Err(err).Msg("failed to clean db")
			}
			log.Info().Msg("cleaned db")

			if err := db.Populate(ctx.Context, sqlite, queries, cheeses); err != nil {
				log.Fatal().Err(err).Msg("failed to populate db")
			}
			log.Info().Msg("populated db")

		}()
		wg.Wait()

		listener, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to listen")
		}

		server := grpc.NewServer()

		healthpb.RegisterHealthServer(server, health.NewServer())
		cheesev1.RegisterCheeseAPIServer(server, cheese.New(queries))

		log.Info().Str("port", port).Msg("api server started")

		return server.Serve(listener)
	},
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("app failed")
	}
}
