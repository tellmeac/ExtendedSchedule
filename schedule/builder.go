package schedule

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"reflect"
	"tellmeac/extended-schedule/common/errors"
	"tellmeac/extended-schedule/userconfig"
	"time"
)

type Provider interface {
	GetByGroup(ctx context.Context, id string, from, to time.Time) (Schedule, error)
	GetByTeacher(ctx context.Context, id string, from, to time.Time) (Schedule, error)
}

type ConfigProvider interface {
	GetByEmail(ctx context.Context, email string) (userconfig.UserConfig, error)
}

var Module = fx.Options(fx.Provide(NewBuilder))

// NewBuilder returns new builder for personal schedule.
func NewBuilder(p Provider, c ConfigProvider) Builder {
	return Builder{
		schedule: p,
		config:   c,
	}
}

// Builder provides methods to receive personal sc.
type Builder struct {
	schedule Provider
	config   ConfigProvider
}

func (b Builder) Personal(ctx context.Context, email string, from, to time.Time) (Schedule, error) {
	settings, err := b.config.GetByEmail(ctx, email)
	switch {
	case errors.Is(err, errors.ErrNotFound):
		log.Warn().Str("email", email).Msg("No settings for user, return empty schedule")
		return emptySchedule(from, to), nil
	case err != nil:
		return Schedule{}, fmt.Errorf("failed to get user settings: %w", err)
	}

	var result Schedule
	if settings.Base != nil {
		var err error
		switch base := (settings.Base).(type) {
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
			panic(fmt.Errorf("unknown base type in settings: %q", reflect.TypeOf(base).Name()))
		}
	}

	// TODO: add extended lessons

	// TODO: apply ignore rules

	return result, nil
}

func emptySchedule(from time.Time, to time.Time) Schedule {
	var daysCount = int(to.Sub(from).Hours() / 24)
	days := make([]Day, daysCount)
	currentDate := from
	for i, _ := range days {
		days[i] = Day{
			Date:    currentDate,
			Lessons: nil,
		}
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return Schedule{
		StartDate: from,
		EndDate:   to,
		Days:      days,
	}
}
