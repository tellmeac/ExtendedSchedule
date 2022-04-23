package repository

import (
	"context"
	"tellmeac/extended-schedule/domain/entity"

	"github.com/google/uuid"
)

type IExcludedLessonsRepository interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]entity.ExcludedLesson, error)
	Update(ctx context.Context, userID uuid.UUID, desired []entity.ExcludedLesson) error
}
