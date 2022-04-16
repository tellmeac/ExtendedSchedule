package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/tellmeac/extended-schedule/domain/entity"
)

type IExtendedRepository interface {
	Update(ctx context.Context, userID uuid.UUID, desired []entity.ExtendedLesson) error
}
