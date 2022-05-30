package provider

import (
	"context"
	"time"
)

type LessonInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	TeacherName string `json:"teacherName"`
	Kind        string `json:"lessonKind"`
}

// ILessonInfoProvider предоставляет методы для получения списка предметов вне расписания.
type ILessonInfoProvider interface {
	GetLessons(ctx context.Context, groupID string, start time.Time, end time.Time) ([]LessonInfo, error)
}

func NewLessonInfoProvider(baseSchedule IBaseScheduleProvider) ILessonInfoProvider {
	return &lessonInfoProvider{
		baseSchedule: baseSchedule,
	}
}

type lessonInfoProvider struct {
	baseSchedule IBaseScheduleProvider
}

func (provider *lessonInfoProvider) GetLessons(ctx context.Context, groupID string, start time.Time, end time.Time) ([]LessonInfo, error) {
	schedule, err := provider.baseSchedule.GetByGroupID(ctx, groupID, start, end)
	if err != nil {
		return nil, err
	}

	var lessonMap = make(map[string]LessonInfo, 0)
	for _, day := range schedule {
		for _, lesson := range day.Lessons {
			if _, ok := lessonMap[lesson.ID]; !ok {
				lessonMap[lesson.ID] = LessonInfo{
					ID:          lesson.ID,
					Name:        lesson.Title,
					TeacherName: lesson.Teacher.Name,
					Kind:        lesson.LessonType.String(),
				}
			}
		}
	}

	var result = make([]LessonInfo, 0, len(lessonMap))
	for _, lesson := range lessonMap {
		result = append(result, lesson)
	}
	return result, nil
}
