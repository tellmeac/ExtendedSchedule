package repository

import (
	"context"
	"tellmeac/extended-schedule/domain/entity"

	"github.com/google/uuid"
)

type IExcludedLessonsRepository interface {
	Update(ctx context.Context, userID uuid.UUID, desired []entity.ExcludedLesson) error
}
