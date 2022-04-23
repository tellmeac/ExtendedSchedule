package repository

import (
	"context"
	"tellmeac/extended-schedule/domain/entity"

	"github.com/google/uuid"
)

type IExtendedRepository interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]entity.ExtendedLesson, error)
	Update(ctx context.Context, userID uuid.UUID, desired []entity.ExtendedLesson) error
}
