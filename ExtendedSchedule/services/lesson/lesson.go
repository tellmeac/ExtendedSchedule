package lesson

import (
	"context"
	"go.uber.org/fx"
	"tellmeac/extended-schedule/domain/schedule"
	scheduleservice "tellmeac/extended-schedule/services/schedule"
	"time"
)

var Module = fx.Options(fx.Provide(New))

// Service provides methods to get lesson without day inner context.
type Service interface {
	GetLessons(ctx context.Context, groupID string, start, end time.Time) ([]schedule.Lesson, error)
}

// New creates default Service implementation.
func New(schedule scheduleservice.Service) Service {
	return &service{
		inner: schedule,
	}
}

type service struct {
	inner scheduleservice.Service
}

func (s *service) GetLessons(ctx context.Context, groupID string, start, end time.Time) ([]schedule.Lesson, error) {
	days, err := s.inner.GetByGroup(ctx, groupID, start, end)
	if err != nil {
		return nil, err
	}

	var lessonMap = make(map[string]schedule.Lesson, 0)
	for _, day := range days {
		for _, lesson := range day.Lessons {
			if _, ok := lessonMap[lesson.ID]; !ok {
				lessonMap[lesson.ID] = toCommonLesson(lesson)
			}
		}
	}

	var result = make([]schedule.Lesson, 0, len(lessonMap))
	for _, lesson := range lessonMap {
		result = append(result, lesson)
	}
	return result, nil
}

func toCommonLesson(lesson schedule.LessonInSchedule) schedule.Lesson {
	return schedule.Lesson{
		ID:         lesson.ID,
		Title:      lesson.Title,
		Teacher:    lesson.Teacher,
		LessonType: lesson.LessonType,
	}
}
