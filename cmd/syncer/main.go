package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"tellmeac/extended-schedule/adapters"
	"tellmeac/extended-schedule/tsuclient"
	"time"
)

const uploadInterval = 500 * time.Millisecond

// Sync teachers and groups in application repository.
func main() {
	ctx := context.Background()
	client := adapters.NewEntClient()
	api := tsuclient.NewTypedClient()
	provider := adapters.NewTargetProvider(api)

	teachers, err := provider.Teachers(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to receive all teachers")
	}

	faculties, err := provider.Faculties(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to receive all faculties")
	}

	tx, err := client.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal().Err(err)
	}

	_ = tx.Teacher.Delete().ExecX(ctx)

	for _, t := range teachers {
		_, err := tx.Teacher.Create().
			SetID(t.ID).
			SetName(t.Name).
			Save(ctx)

		if err != nil {
			log.Error().Err(err).Msg("Failed to save teacher")
		}
	}

	_ = tx.StudyGroup.Delete().ExecX(ctx)

	for _, faculty := range faculties {
		groups, err := provider.GroupsByFaculty(ctx, faculty.ID)
		time.Sleep(uploadInterval)

		if err != nil {
			log.Error().Err(err).Str("facultyName", faculty.Name).
				Str("facultyID", faculty.ID).
				Msg("Failed to receive groups by faculty")
			continue
		}

		for _, g := range groups {
			_, err := tx.StudyGroup.Create().
				SetID(g.ID).
				SetName(g.Name).
				SetFacultyName(faculty.Name).
				Save(ctx)

			if err != nil {
				log.Error().Err(err).Msg("Failed to save group")
			}
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatal().Err(err)
	}

	log.Info().Msg("Syncer is done")
}
