package ent

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"go.uber.org/fx"
	appconfig "tellmeac/extended-schedule/pkg/config"

	// Required to connect to postgres database
	_ "github.com/jackc/pgx/v4/stdlib"
)

var Module = fx.Options(fx.Provide(NewEntClient))

// NewEntClient returns new ent client.
func NewEntClient(cfg appconfig.Config) *Client {
	db, err := sql.Open("pgx", cfg.Database.ConnectionAddress)
	if err != nil {
		panic(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := NewClient(Driver(drv))

	if err = client.Schema.Create(context.Background()); err != nil {
		panic(err)
	}
	return client
}
