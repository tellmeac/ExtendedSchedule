package schedule

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/tellmeac/ext-schedule/schedule/common/userconfig"
	"reflect"
	"time"
)

type Provider interface {
	GetByGroup(ctx context.Context, id string, from, to time.Time) (Schedule, error)
	GetByTeacher(ctx context.Context, id string, from, to time.Time) (Schedule, error)
}

type ConfigProvider interface {
	Get(ctx context.Context, userID uuid.UUID) (userconfig.UserConfig, error)
}

// Builder provides methods to receive personal schedule.
type Builder struct {
	schedule Provider
	config   ConfigProvider
}

func (b Builder) Personal(ctx context.Context, userID uuid.UUID, from, to time.Time) (Schedule, error) {
	settings, err := b.config.Get(ctx, userID)
	if err != nil {
		return Schedule{}, fmt.Errorf("failed to get user settings: %w", err)
	}

	var result Schedule
	if settings.Base != nil {
		var err error
		switch base := (*settings.Base).(type) {
		case userconfig.Teacher:
			result, err = b.schedule.GetByTeacher(ctx, base.ID, from, to)
			if err != nil {
				return Schedule{}, fmt.Errorf("failed to get teacher schedule: %w", err)
			}
		case userconfig.StudyGroup:
			result, err = b.schedule.GetByGroup(ctx, base.ID, from, to)
			if err != nil {
				return Schedule{}, fmt.Errorf("failed to get base group schedule: %w", err)
			}
		default:
			panic(fmt.Errorf("unknown base in settings: %q", reflect.TypeOf(base).Name()))
		}
	}

	// TODO: extended lessons

	// TODO: apply ignore rules

	return result, nil
}
