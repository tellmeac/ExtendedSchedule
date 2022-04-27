package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"tellmeac/extended-schedule/adapters/ent"
	"tellmeac/extended-schedule/adapters/ent/excludedlesson"
	"tellmeac/extended-schedule/adapters/ent/joinedgroups"
	"tellmeac/extended-schedule/domain/aggregates"
	"tellmeac/extended-schedule/domain/entity"
)

type EntUserConfigRepository struct {
	client *ent.Client
}

func (repository EntUserConfigRepository) Get(ctx context.Context, userID uuid.UUID) (aggregates.UserConfig, error) {
	tx, err := repository.client.Tx(ctx)
	if err != nil {
		return aggregates.UserConfig{}, err
	}

	joinedGroups, err := tx.JoinedGroups.Query().Where(joinedgroups.UserIDEQ(userID)).Only(ctx)
	if err != nil {
		return aggregates.UserConfig{}, rollback(tx, fmt.Errorf("failed to get joined groups: %w", err))
	}

	excluded, err := tx.ExcludedLesson.Query().Where(excludedlesson.UserIDEQ(userID)).All(ctx)
	if err != nil {
		return aggregates.UserConfig{}, rollback(tx, fmt.Errorf("failed to get excluded lessons: %w", err))
	}

	if err := tx.Commit(); err != nil {
		return aggregates.UserConfig{}, fmt.Errorf("failed to commit select request: %w", err)
	}

	return aggregates.UserConfig{
		UserID:          userID,
		JoinedGroups:    joinedGroups.JoinedGroups,
		ExcludedLessons: mapExcluded(excluded),
	}, nil
}

func mapExcluded(excluded []*ent.ExcludedLesson) []entity.ExcludedLesson {
	var result = make([]entity.ExcludedLesson, 0, len(excluded))
	for _, e := range excluded {
		result = append(result, entity.ExcludedLesson{
			ID:       e.ID,
			LessonID: e.LessonID,
			Teacher:  e.Teacher,
			Position: e.Position,
			WeekDay:  e.Weekday,
		})
	}
	return result
}

func (repository EntUserConfigRepository) Update(ctx context.Context, userID uuid.UUID, desired aggregates.UserConfig) error {
	//TODO implement me
	panic("implement me")
}

func rollback(tx *ent.Tx, err error) error {
	if rollbackErr := tx.Rollback(); rollbackErr != nil {
		err = fmt.Errorf("%w: %v", err, rollbackErr)
	}
	return err
}
