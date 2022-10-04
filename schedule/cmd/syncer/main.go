package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"tellmeac/extended-schedule/adapters"
	"tellmeac/extended-schedule/common/tsu"
)

// Sync teachers and groups in application repository.
func main() {
	ctx := context.Background()
	client := adapters.NewEntClient()
	api := tsu.NewTypedClient()
	provider := adapters.NewTargetProvider(api)

	teachers, err := provider.Teachers(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to receive all teachers")
	}

	tx, err := client.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal().Err(err)
	}

	_ = tx.Teacher.Delete().ExecX(ctx)

	for _, t := range teachers {
		tx.Teacher.Create().
			SetID(t.ID).
			SetName(t.Name).
			SaveX(ctx)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal().Err(err)
	}
}
