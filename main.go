package main

import (
	"fmt"
	"net"
	"os"

	"github.com/fredouric/cheese-finder-grpc/api/cheese"
	"github.com/fredouric/cheese-finder-grpc/dataset"
	"github.com/fredouric/cheese-finder-grpc/pb/cheesev1"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

var (
	port string
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
	},
	Action: func(ctx *cli.Context) error {

		log.Level(zerolog.DebugLevel)
		go func() {
			cheeses, err := dataset.Fetch()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to fetch cheeses")
			}
			for _, cheese := range cheeses {
				fmt.Println(cheese.Fromage, cheese.Lait, cheese.Departement)
			}
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
