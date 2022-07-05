package infrastructure

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"tellmeac/extended-schedule/config"
	"tellmeac/extended-schedule/infrastructure/ent"

	// Required to connect to postgres database
	_ "github.com/jackc/pgx/v4/stdlib"
)

// NewEntClient returns new ent client.
func NewEntClient(cfg config.Config) *ent.Client {
	db, err := sql.Open("pgx", cfg.Database.ConnectionAddress)
	if err != nil {
		panic(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	if err = client.Schema.Create(context.Background()); err != nil {
		panic(err)
	}
	return client
}
