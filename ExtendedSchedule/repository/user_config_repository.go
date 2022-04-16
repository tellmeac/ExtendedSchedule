package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/tellmeac/extended-schedule/domain/aggregates"
)

type IUserConfigRepository interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) (aggregates.UserConfig, error)
}
