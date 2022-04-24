package factory

import (
	"context"
	"errors"
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

	schedule, err = factory.makeWithExtendedLessons(ctx, userID, start, end, schedule)
	if err != nil {
		return nil, fmt.Errorf("failed to build schedule with extended lessons: %w", err)
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
	schedule []entity.DaySchedule,
) ([]entity.DaySchedule, error) {
	groups, err := factory.joinedGroups.GetByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user joined groups: %w", err)
	}

	for _, group := range groups {
		scheduleByGroup, err := factory.scheduleProvider.GetByGroupID(ctx, group.ID, start, end)
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
	lessons, err := factory.extendedLessons.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	for _, lesson := range lessons {
		extendedSchedule, err := factory.scheduleProvider.GetLessonSchedule(
			ctx,
			lesson.Ref.GroupID,
			lesson.Ref.LessonID,
			start,
			end,
		)
		if err != nil {
			return nil, err
		}

		schedule, err = joinSchedules(schedule, extendedSchedule)
		if err != nil {
			return nil, err
		}
	}

	return schedule, nil
}

// joinSchedules объединяет два одинаковых по размеру и датам списка расписания.
func joinSchedules(a []entity.DaySchedule, b []entity.DaySchedule) ([]entity.DaySchedule, error) {
	if a == nil {
		return b, nil
	}
	if b == nil {
		return a, nil
	}

	if len(a) != len(b) {
		return nil, errors.New("expected to have equal schedule length")
	}

	var joinedResult = a
	for i := 0; i < len(a); i++ {
		if a[i].Date != b[i].Date {
			return nil, errors.New("expected to have equal date of day index by index")
		}
		joinedResult[i].Sections = joinSections(joinedResult[i].Sections, b[i].Sections)
	}

	return joinedResult, nil
}

func joinSections(x []entity.Section, y []entity.Section) []entity.Section {
	var joinedResult = x
	for i := 0; i < len(x); i++ {
		joinedResult[i].Lessons = joinLessons(joinedResult[i].Lessons, y[i].Lessons)
	}

	return joinedResult
}

type lessonJoinKey struct {
	ID           string
	LessonType   string
	LessonNumber int
}

func joinLessons(g []entity.Lesson, h []entity.Lesson) []entity.Lesson {
	var joinMap = make(map[lessonJoinKey]*entity.Lesson)
	var key lessonJoinKey
	for i := 0; i < len(g); i++ {
		key = lessonJoinKey{
			ID:           g[i].ID,
			LessonType:   g[i].LessonType,
			LessonNumber: g[i].LessonNumber,
		}
		joinMap[key] = &g[i]
	}
	for i := 0; i < len(h); i++ {
		key = lessonJoinKey{
			ID:           h[i].ID,
			LessonType:   h[i].LessonType,
			LessonNumber: h[i].LessonNumber,
		}
		joinMap[key] = &h[i]
	}

	var joinedResult = make([]entity.Lesson, 0, len(joinMap))
	for _, lesson := range joinMap {
		joinedResult = append(joinedResult, *lesson)
	}

	return joinedResult
}
