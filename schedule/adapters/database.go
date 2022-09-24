package adapters

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/rs/zerolog/log"
	"tellmeac/extended-schedule/adapters/ent"
	"tellmeac/extended-schedule/config"

	// Required to connect to postgres database
	_ "github.com/jackc/pgx/v4/stdlib"
)

// NewEntClient returns new ent client.
func NewEntClient() *ent.Client {
	cfg := config.Get()

	db, err := sql.Open("pgx", cfg.DatabaseAddress)
	if err != nil {
		panic(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	if err = client.Schema.Create(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("Failed to create schema in database")
	}
	return client
}
