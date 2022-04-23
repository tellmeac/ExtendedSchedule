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
	groups, err := factory.joinedGroups.GetByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user joined groups: %w", err)
	}

	var schedule = make([]entity.DaySchedule, 0)
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
