package lesson

import (
	"context"
	"github.com/tellmeac/extended-schedule/userconfig/domain/lesson"
	"github.com/tellmeac/extended-schedule/userconfig/domain/schedule"
	scheduleservice "github.com/tellmeac/extended-schedule/userconfig/services/schedule"
	"time"
)

// Service provides methods to get lesson without day inner context.
type Service interface {
	GetLessons(ctx context.Context, groupID string, start, end time.Time) ([]lesson.Lesson, error)
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

func (s *service) GetLessons(ctx context.Context, groupID string, start, end time.Time) ([]lesson.Lesson, error) {
	days, err := s.inner.GetByGroup(ctx, groupID, start, end)
	if err != nil {
		return nil, err
	}

	var lessonMap = make(map[string]schedule.Lesson, 0)
	for _, day := range days {
		for _, l := range day.Lessons {
			if _, ok := lessonMap[l.ID]; !ok {
				lessonMap[l.ID] = l
			}
		}
	}

	var result = make([]schedule.Lesson, 0, len(lessonMap))
	for _, l := range lessonMap {
		result = append(result, l)
	}

	return toLessons(result), nil
}

func toLessons(ctxLessons []schedule.Lesson) []lesson.Lesson {
	var result = make([]lesson.Lesson, len(ctxLessons))
	for i, l := range ctxLessons {
		result[i] = l.Lesson
	}
	return result
}
