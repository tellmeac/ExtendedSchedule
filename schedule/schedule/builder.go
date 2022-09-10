package schedule

import (
	"context"
	"github.com/google/uuid"
	"github.com/tellmeac/ext-schedule/schedule/common/userconfig"
	"time"
)

type Provider interface {
	GetByGroup(ctx context.Context, id string, from, to time.Time) (Schedule, error)
}

type ConfigProvider interface {
	Get(ctx context.Context, userID uuid.UUID) (userconfig.UserConfig, error)
}

type Builder struct {
	schedule Provider
	config   ConfigProvider
}

func (b Builder) Personal(ctx context.Context, userID uuid.UUID, from, to time.Time) (Schedule, error) {
	return Schedule{}, nil
}
