package db

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
)

var (
	//go:embed migrations
	embedMigrations embed.FS
)

func Migrate(dbPath string) *Queries {
	goose.SetLogger(goose.NopLogger())

	sqlite, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open db")
	}

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
