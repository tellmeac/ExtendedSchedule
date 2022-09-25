//go:build ignore

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
	// Create a local migration directory able to understand Atlas migration file format for replay.
	dir, err := atlas.NewLocalDir("migrations")
	if err != nil {
		log.Fatalf("Failed to init atlas migration directory: %v", err)
	}

	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.Postgres),        // Ent dialect to use
		schema.WithFormatter(atlas.DefaultFormatter),
	}

	if len(os.Args) != 2 {
		log.Fatalln("Migration name is required. Use: 'go run -mod=mod migrations.go <name>'")
	}

	addr := flag.String("url", "postgres://postgres:postgres@localhost:5432/ExtendedSchedule?sslmode=disable", "")
	flag.Parse()

	// Generate migrations using Atlas
	err = migrate.NamedDiff(ctx, addr, os.Args[1], opts...)
	if err != nil {
		log.Fatalf("Failed generating migration file: %v", err)
	}
}
