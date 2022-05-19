package infrastructure

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"tellmeac/extended-schedule/adapters/ent"
	"tellmeac/extended-schedule/config"
	"tellmeac/extended-schedule/infrastructure/log"

	// Required to connect to postgres database
	_ "github.com/jackc/pgx/v4/stdlib"
)

// NewClient returns new ent client from config values.
func NewClient(cfg config.Config) *ent.Client {
	db, err := sql.Open("pgx", cfg.Database.ConnectionAddress)
	if err != nil {
		log.Sugared.Error("Failed to setup database connection", "error", err)
		panic(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	log.Sugared.Info("Put database schema")
	if err = client.Schema.Create(context.Background()); err != nil {
		log.Sugared.Error("Failed to create database schema", "error", err)
		panic(err)
	}
	return client
}
