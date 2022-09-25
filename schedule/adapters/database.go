package adapters

import (
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/rs/zerolog/log"
	"tellmeac/extended-schedule/adapters/ent"
	"tellmeac/extended-schedule/config"

	// Required to connect to postgres database
	_ "github.com/lib/pq"
)

// NewEntClient returns new ent client.
func NewEntClient() *ent.Client {
	cfg := config.Get()

	db, err := sql.Open("postgres", cfg.DatabaseAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open connection to database")
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	if cfg.Debug {
		client = client.Debug()
	}

	return client
}
