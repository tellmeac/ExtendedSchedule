package main

import (
	"context"
	"flag"
	"log"
	"os"

	"tellmeac/extended-schedule/adapters/ent/migrate"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	dir, err := atlas.NewLocalDir("migrations")
	if err != nil {
		log.Fatalf("Failed to init atlas migration directory: %v", err)
	}

	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),
		schema.WithMigrationMode(schema.ModeReplay),
		schema.WithDialect(dialect.Postgres),
		schema.WithFormatter(atlas.DefaultFormatter),
	}

	if len(os.Args) != 2 {
		log.Fatalln("Migration name is required. Use: 'go run -mod=mod migrations.go <name>'")
	}

	addr := flag.String("url", "postgres://postgres:postgres@localhost:5432/ExtendedSchedule?sslmode=disable", "")
	flag.Parse()

	// Generate migrations using Atlas
	err = migrate.NamedDiff(ctx, *addr, os.Args[1], opts...)
	if err != nil {
		log.Fatalf("Failed generating migration file: %v", err)
	}
}
