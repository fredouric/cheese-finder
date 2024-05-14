package main

import (
	_ "embed"
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
		},
		&cli.StringFlag{
			Name:        "DB Path",
			Value:       "./cheese.sqlite3",
			Destination: &dbPath,
		},
	},
	Action: func(ctx *cli.Context) error {
		log.Level(zerolog.DebugLevel)

		queries := db.Migrate(dbPath)

		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			cheeses, err := dataset.Fetch()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to fetch dataset")
			}

			log.Info().Msg("entire dataset fetched")

			// TODO: populate db with all cheeeses
			// TODO: mappers for structs
		}()

		wg.Wait()

		listener, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to listen")
		}

		server := grpc.NewServer()

		cheesev1.RegisterCheeseAPIServer(server, cheese.New())

		log.Info().Str("port", port).Msg("api server started")

		return server.Serve(listener)
	},
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("app failed")
	}
}
