package adapters

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"tellmeac/extended-schedule/pkg/tsuclient"
	"tellmeac/extended-schedule/schedule"
)

func NewScheduleProvider(client tsuclient.ClientWithResponsesInterface) ScheduleProvider {
	return ScheduleProvider{client: client}
}

type ScheduleProvider struct {
	client tsuclient.ClientWithResponsesInterface
}

func (sp ScheduleProvider) GetByTeacher(ctx context.Context, id string, from, to time.Time) (schedule.Schedule, error) {
	resp, err := sp.client.GetScheduleProfessorWithResponse(ctx, &tsuclient.GetScheduleProfessorParams{
		ID:       id,
		DateFrom: openapi_types.Date{Time: from},
		DateTo:   openapi_types.Date{Time: to},
	})
	switch {
	case err != nil:
		return schedule.Schedule{}, fmt.Errorf("failed to get response: %w", err)
	case resp.StatusCode() != 200:
		return schedule.Schedule{}, fmt.Errorf("failed with status code = %d", resp.StatusCode())
	}

	return fromScheduleDto(*resp.JSON200, from, to)
}

func (sp ScheduleProvider) GetByGroup(ctx context.Context, id string, from, to time.Time) (schedule.Schedule, error) {
	resp, err := sp.client.GetScheduleGroupWithResponse(ctx, &tsuclient.GetScheduleGroupParams{
		ID:       id,
		DateFrom: openapi_types.Date{Time: from},
		DateTo:   openapi_types.Date{Time: to},
	})
	switch {
	case err != nil:
		return schedule.Schedule{}, fmt.Errorf("failed to get response: %w", err)
	case resp.HTTPResponse.StatusCode != 200:
		return schedule.Schedule{}, fmt.Errorf("failed with status code = %d", resp.HTTPResponse.StatusCode)
	}

	return fromScheduleDto(*resp.JSON200, from, to)
}

func fromScheduleDto(days []tsuclient.DaySchedule, from, to time.Time) (schedule.Schedule, error) {
	s := schedule.Schedule{
		StartDate: from,
		EndDate:   to,
	}

	if days == nil {
		return s, nil
	}
	s.Days = make([]schedule.Day, 0, len(days))

	for _, d := range days {
		s.Days = append(s.Days, schedule.Day{
			Date:    d.Date.Time,
			Lessons: fromLessonsDto(d.Lessons),
		})
	}

	return s, nil
}

func fromLessonsDto(lessons []tsuclient.Lesson) []schedule.Lesson {
	slices.SortFunc(lessons, func(a, b tsuclient.Lesson) bool {
		return a.Position < b.Position
	})

	result := make([]schedule.Lesson, 0)
	for _, l := range lessons {
		if l.Type != "LESSON" {
			continue
		}

		// if teacher is not set it has no identifier.
		var teacher *schedule.Teacher
		if l.Professor.ID != nil {
			teacher = &schedule.Teacher{
				ID:   *l.Professor.ID,
				Name: *l.Professor.FullName,
			}
		}

		lesson := schedule.Lesson{
			ID:      *l.ID,
			Kind:    *l.LessonKind,
			Name:    *l.Title,
			Pos:     l.Position,
			Teacher: teacher,
			Groups: lo.Map(*l.Groups, func(g tsuclient.StudyGroup, _ int) string {
				return g.Name
			}),
		}

		result = append(result, lesson)
	}

	return result
}
