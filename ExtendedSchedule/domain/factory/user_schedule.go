package factory

import (
	"context"
	"fmt"
	"tellmeac/extended-schedule/domain/aggregates"
	"tellmeac/extended-schedule/domain/providers"
	"tellmeac/extended-schedule/domain/repository"
	"time"

	"github.com/google/uuid"
)

type IUserScheduleFactory interface {
	Make(ctx context.Context, userID uuid.UUID, start time.Time, end time.Time) ([]aggregates.DaySchedule, error)
}

type UsersScheduleFactory struct {
	scheduleProvider providers.IBaseScheduleProvider
	joinedGroups     repository.IJoinedGroupsRepository
	excludedLessons  repository.IExcludedLessonsRepository
}

func (factory *UsersScheduleFactory) Make(ctx context.Context, userID uuid.UUID, start time.Time, end time.Time) ([]aggregates.DaySchedule, error) {
	var schedule []aggregates.DaySchedule
	var err error

	schedule, err = factory.makeWithJoinedGroups(ctx, userID, start, end, schedule)
	if err != nil {
		return nil, fmt.Errorf("failed to join user schedule by groups: %w", err)
	}

	schedule, err = factory.makeWithExcludedLessons(ctx, userID, start, end, schedule)
	if err != nil {
		return nil, fmt.Errorf("failed to cut schedule with excluded lessons: %w", err)
	}

	return schedule, nil
}

func (factory *UsersScheduleFactory) makeWithJoinedGroups(
	ctx context.Context,
	userID uuid.UUID,
	start time.Time,
	end time.Time,
	schedule []aggregates.DaySchedule,
) ([]aggregates.DaySchedule, error) {
	groups, err := factory.joinedGroups.GetByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user joined groups: %w", err)
	}

	for _, group := range groups {
		scheduleByGroup, err := factory.scheduleProvider.GetByGroupID(ctx, group.ID, start, end)
		if err != nil {
			return nil, fmt.Errorf("failed to get schedule by group: %w", err)
		}

		schedule, err = aggregates.JoinSchedules(schedule, scheduleByGroup)
		if err != nil {
			return nil, fmt.Errorf("failed to join schedule in one: %w", err)
		}
	}

	return schedule, nil
}

func (factory *UsersScheduleFactory) makeWithExcludedLessons(
	ctx context.Context,
	userID uuid.UUID,
	_ time.Time,
	_ time.Time,
	schedule []aggregates.DaySchedule,
) ([]aggregates.DaySchedule, error) {
	excludedLessons, err := factory.excludedLessons.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	for _, day := range schedule {
		if err := day.ExcludeLessons(excludedLessons); err != nil {
			return nil, err
		}
	}

	return schedule, nil
}
