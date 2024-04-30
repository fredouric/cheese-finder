package main

import (
	"database/sql"
	"net"
	"os"

	"github.com/fredouric/cheese-finder-grpc/api/cheese"
	"github.com/fredouric/cheese-finder-grpc/dataset"
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
)

var app = &cli.App{
	Name:    "cheese-finder-grpc",
	Usage:   "",
	Suggest: true,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "port",
			Value:       ":8080",
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

		db, err := sql.Open("sqlite", dbPath)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to open db")
		}

		go func() {
			cheeses, err := dataset.Fetch()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to fetch cheeses")
			}

			// Populate db
		}()

		listener, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to listen")
		}

		server := grpc.NewServer()

		cheesev1.RegisterCheeseAPIServer(server, cheese.New())

		return server.Serve(listener)
	},
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("app failed")
	}
}
