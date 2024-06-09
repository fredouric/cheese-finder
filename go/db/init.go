package db

import (
	"context"
	"database/sql"
	"embed"

	"github.com/fredouric/cheese-finder-grpc/dataset"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
)

var (
	//go:embed migrations
	embedMigrations embed.FS
)

func Migrate(sqlite *sql.DB, dbPath string) *Queries {
	goose.SetLogger(goose.NopLogger())

	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("sqlite"); err != nil {
		log.Fatal().Err(err).Msg("failed to set migration dialect")
	}

	if err := goose.Up(sqlite, "migrations"); err != nil {
		log.Fatal().Err(err).Msg("failed to migrate db")
	}
	log.Info().Msg("migrated db")

	return New(sqlite)
}

func Populate(ctx context.Context, sqlite *sql.DB, queries *Queries, cheeses []dataset.Cheese) error {
	tx, err := sqlite.Begin()
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	qtx := queries.WithTx(tx)
	for _, cheese := range cheeses {
		mappedCheese := CheeseMapper(cheese)
		if err := qtx.AddCheese(ctx, AddCheeseParams{
			Departement:   mappedCheese.Departement,
			Fromage:       mappedCheese.Fromage,
			Pagefrancaise: mappedCheese.Pagefrancaise,
			Englishpage:   mappedCheese.Englishpage,
			Lait:          mappedCheese.Lait,
			Geoshape:      mappedCheese.Geoshape,
			Geopoint2d:    mappedCheese.Geopoint2d,
		}); err != nil {
			return err
		}
	}
	return tx.Commit()
}
