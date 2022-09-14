package adapters

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/tellmeac/ext-schedule/schedule/common/tsu"
	"github.com/tellmeac/ext-schedule/schedule/schedule"
)

func NewScheduleProvider(client tsu.ClientWithResponsesInterface) ScheduleProvider {
	return ScheduleProvider{client: client}
}

type ScheduleProvider struct {
	client tsu.ClientWithResponsesInterface
}

func (sp ScheduleProvider) GetByTeacher(ctx context.Context, id string, from, to time.Time) (schedule.Schedule, error) {
	resp, err := sp.client.GetScheduleProfessorWithResponse(ctx, &tsu.GetScheduleProfessorParams{
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

func (sp ScheduleProvider) GetByGroup(ctx context.Context, id string, from, to time.Time) (schedule.Schedule, error) {
	resp, err := sp.client.GetScheduleGroupWithResponse(ctx, &tsu.GetScheduleGroupParams{
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

func fromScheduleDto(days []tsu.DaySchedule, from, to time.Time) (schedule.Schedule, error) {
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
			Date:  d.Date.Time,
			Cells: groupByCells(d.Lessons),
		})
	}

	return s, nil
}

func groupByCells(lessons []tsu.Lesson) []schedule.Cell {
	slices.SortFunc(lessons, func(a, b tsu.Lesson) bool {
		return a.Position < b.Position
	})

	const maxCells = 7
	result := make([]schedule.Cell, 0, maxCells)
	currentCell := schedule.Cell{Pos: 0}
	for _, l := range lessons {
		if l.Position > currentCell.Pos {
			result = append(result, currentCell)
			currentCell = schedule.Cell{
				Pos:     l.Position,
				Lessons: nil,
			}
		}

		// type can be "EMPTY"
		if l.Type == "LESSON" {
			lesson := schedule.Lesson{
				ID:   *l.ID,
				Kind: *l.LessonKind,
				Name: *l.Title,
				Groups: lo.Map(*l.Groups, func(g tsu.StudyGroup, _ int) string {
					return g.Name
				}),
			}
			// if teacher is not set it has no identifier.
			if l.Professor.ID != nil {
				lesson.Teacher = &schedule.Teacher{
					ID:   *l.Professor.ID,
					Name: *l.Professor.FullName,
				}
			}

			currentCell.Lessons = append(currentCell.Lessons, lesson)
		}
	}
	return result
}
