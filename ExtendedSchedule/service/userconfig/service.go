package userconfig

import (
	"context"
	"github.com/google/uuid"
	"tellmeac/extended-schedule/domain/aggregates"
)

type IService interface {
	GetUserConfig(ctx context.Context, userID uuid.UUID) (aggregates.UserConfig, error)
	UpdateUserConfig(ctx context.Context, userID uuid.UUID, desired aggregates.UserConfig) error
}
