package main

import (
	"database/sql"
	_ "embed"
	"net"
	"os"

	"github.com/fredouric/cheese-finder-grpc/api/cheese"
	"github.com/fredouric/cheese-finder-grpc/dataset"
	"github.com/fredouric/cheese-finder-grpc/db"
	"github.com/fredouric/cheese-finder-grpc/pb/cheesev1"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

var (
	port   string
	dbPath string
	//go:embed db/schema.sql
	schema string
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

		sqlite, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to open db")
		}

		if _, err := sqlite.ExecContext(ctx.Context, schema); err != nil {
			log.Fatal().Err(err).Msg("failed to create tables")
		}

		queries := db.New(sqlite)

		go func() {
			cheeses, err := dataset.Fetch()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to fetch cheeses")
			}
			log.Info().Int("totalCheeses", len(cheeses)).Msg("fetched dataset")

			if err := queries.AddCheese(ctx.Context, db.AddCheeseParams{
				Departement:   "Some Departement",
				Fromage:       "Some Fromage",
				Pagefrancaise: "Some Page Francaise",
				Englishpage:   "Some English Page",
				Lait:          "Lait, Brebis",
				Geoshape:      "{}",          // Placeholder for GeoShape data
				Geopoint2d:    "1.234,5.678", // Placeholder for GeoPoint2D data
			}); err != nil {
				log.Fatal().Err(err).Msg("failed to populate db")
			}
		}()

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
