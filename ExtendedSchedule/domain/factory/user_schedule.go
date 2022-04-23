package factory

import (
	"context"
	"fmt"
	"tellmeac/extended-schedule/domain/entity"
	"tellmeac/extended-schedule/domain/providers"
	"tellmeac/extended-schedule/domain/repository"
	"time"

	"github.com/google/uuid"
)

type IUserScheduleFactory interface {
	Make(ctx context.Context, userID uuid.UUID, start time.Time, end time.Time) ([]entity.DaySchedule, error)
}

type UsersScheduleFactory struct {
	scheduleProvider providers.IBaseScheduleProvider
	joinedGroups     repository.IJoinedGroupsRepository
	extendedLessons  repository.IExtendedRepository
	excludedLessons  repository.IExcludedLessonsRepository
}

func (factory *UsersScheduleFactory) Make(ctx context.Context, userID uuid.UUID, start time.Time, end time.Time) ([]entity.DaySchedule, error) {
	var schedule []entity.DaySchedule
	var err error

	schedule, err = factory.makeWithJoinedGroups(ctx, userID, start, end, schedule)
	if err != nil {
		return nil, fmt.Errorf("failed to join user schedule by groups: %w", err)
	}

	schedule, err = factory.makeWithExcludedLessons(ctx, userID, start, end, schedule)
	if err != nil {
		return nil, fmt.Errorf("failed to cut schedule with excluded lessons: %w", err)
	}

	schedule, err = factory.makeWithExtendedLessons(ctx, userID, start, end, schedule)
	if err != nil {
		return nil, fmt.Errorf("failed to build schedule with extended lessons: %w", err)
	}

	return schedule, nil
}

func (factory *UsersScheduleFactory) makeWithJoinedGroups(
	ctx context.Context,
	userID uuid.UUID,
	start time.Time,
	end time.Time,
	schedule []entity.DaySchedule,
) ([]entity.DaySchedule, error) {
	groups, err := factory.joinedGroups.GetByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user joined groups: %w", err)
	}

	for _, group := range groups {
		scheduleByGroup, err := factory.scheduleProvider.GetByGroup(ctx, group.ID, start, end)
		if err != nil {
			return nil, fmt.Errorf("failed to get schedule by group: %w", err)
		}

		schedule, err = joinSchedules(schedule, scheduleByGroup)
		if err != nil {
			return nil, fmt.Errorf("failed to join schedule in one: %w", err)
		}
	}

	return schedule, nil
}

func (factory *UsersScheduleFactory) makeWithExcludedLessons(
	ctx context.Context,
	userID uuid.UUID,
	start time.Time,
	end time.Time,
	schedule []entity.DaySchedule,
) ([]entity.DaySchedule, error) {
	return schedule, nil
}

func (factory *UsersScheduleFactory) makeWithExtendedLessons(
	ctx context.Context,
	userID uuid.UUID,
	start time.Time,
	end time.Time,
	schedule []entity.DaySchedule,
) ([]entity.DaySchedule, error) {
	return schedule, nil
}

// joinSchedules соединяет два списка расписания.
func joinSchedules(a []entity.DaySchedule, b []entity.DaySchedule) ([]entity.DaySchedule, error) {
	if a == nil {
		return b, nil
	}
	if b == nil {
		return a, nil
	}

	return nil, nil
}
