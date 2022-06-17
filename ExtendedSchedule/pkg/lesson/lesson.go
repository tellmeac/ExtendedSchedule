package lesson

import (
	"context"
	"go.uber.org/fx"
	commonmodels "tellmeac/extended-schedule/common/models"
	"tellmeac/extended-schedule/pkg/schedule"
	"time"
)

var Module = fx.Options(fx.Provide(New))

// Manager provides methods to get lesson without day schedule context.
type Manager interface {
	GetLessons(ctx context.Context, groupID string, start time.Time, end time.Time) ([]commonmodels.Lesson, error)
}

// New creates default Manager implementation.
func New(schedule schedule.Manager) Manager {
	return &manager{
		schedule: schedule,
	}
}

type manager struct {
	schedule schedule.Manager
}

func (m *manager) GetLessons(ctx context.Context, groupID string, start time.Time, end time.Time) ([]commonmodels.Lesson, error) {
	days, err := m.schedule.GetByGroup(ctx, groupID, start, end)
	if err != nil {
		return nil, err
	}

	var lessonMap = make(map[string]commonmodels.Lesson, 0)
	for _, day := range days {
		for _, lesson := range day.Lessons {
			if _, ok := lessonMap[lesson.ID]; !ok {
				lessonMap[lesson.ID] = toCommonLesson(lesson)
			}
		}
	}

	var result = make([]commonmodels.Lesson, 0, len(lessonMap))
	for _, lesson := range lessonMap {
		result = append(result, lesson)
	}
	return result, nil
}

func toCommonLesson(lesson commonmodels.LessonWithContext) commonmodels.Lesson {
	return commonmodels.Lesson{
		ID:         lesson.ID,
		Title:      lesson.Title,
		Teacher:    lesson.Teacher,
		LessonType: lesson.LessonType,
	}
}
