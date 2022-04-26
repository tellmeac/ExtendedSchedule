package providers

import (
	"context"
	"fmt"
	"tellmeac/extended-schedule/adapters/clients/tsuschedule"
	"tellmeac/extended-schedule/domain/aggregates"
	"tellmeac/extended-schedule/domain/entity"
	"tellmeac/extended-schedule/domain/values"
	"time"
)

type BaseScheduleProvider struct {
	client *tsuschedule.Client
}

func (provider *BaseScheduleProvider) GetLessonSchedule(ctx context.Context, groupID string, lessonID string, start time.Time, end time.Time) ([]aggregates.DaySchedule, error) {
	params := tsuschedule.GetScheduleGroupParams{
		Id:       groupID,
		DateFrom: start.Format("2006-01-02"),
		DateTo:   end.Format("2006-01-02"),
	}

	response, err := provider.client.GetScheduleGroup(ctx, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to get response from api: %w", err)
	}

	scheduleDto, err := tsuschedule.ParseGetScheduleGroupResponse(response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response from api: %w", err)
	}

	if scheduleDto.JSON200 == nil {
		return nil, fmt.Errorf("failed to get schedule from parsed response: %w", err)
	}

	var result = make([]aggregates.DaySchedule, len(*scheduleDto.JSON200))
	for i, day := range *scheduleDto.JSON200 {
		result[i], err = mapDaySchedule(day)
		if err != nil {
			return nil, fmt.Errorf("failed to map response data properly: %w", err)
		}
		result[i] = filterByLessonID(result[i], lessonID)
	}

	return result, nil
}

func filterByLessonID(day aggregates.DaySchedule, lessonID string) aggregates.DaySchedule {
	var filteredLessons = make([]entity.Lesson, 0)
	for _, lesson := range day.Lessons {
		if lesson.ID == lessonID {
			filteredLessons = append(filteredLessons, lesson)
		}
	}
	return day
}

func (provider *BaseScheduleProvider) GetByGroupID(
	ctx context.Context,
	groupID string,
	start time.Time,
	end time.Time,
) ([]aggregates.DaySchedule, error) {
	params := tsuschedule.GetScheduleGroupParams{
		Id:       groupID,
		DateFrom: start.Format("2006-01-02"),
		DateTo:   end.Format("2006-01-02"),
	}

	response, err := provider.client.GetScheduleGroup(ctx, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to get response from api: %w", err)
	}

	scheduleDto, err := tsuschedule.ParseGetScheduleGroupResponse(response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response from api: %w", err)
	}

	if scheduleDto.JSON200 == nil {
		return nil, fmt.Errorf("failed to get schedule from parsed response: %w", err)
	}

	var result = make([]aggregates.DaySchedule, len(*scheduleDto.JSON200))
	for i, day := range *scheduleDto.JSON200 {
		result[i], err = mapDaySchedule(day)
		if err != nil {
			return nil, fmt.Errorf("failed to map response data properly: %w", err)
		}
	}

	return result, nil
}

func mapDaySchedule(day tsuschedule.DaySchedule) (aggregates.DaySchedule, error) {
	date, err := time.Parse("2006-01-02", day.Date)
	if err != nil {
		return aggregates.DaySchedule{}, err
	}

	scheduleDay := aggregates.DaySchedule{
		Date:    date,
		Lessons: make([]entity.Lesson, 0, len(day.Lessons)),
	}

	for _, lesson := range day.Lessons {
		// ignoring empty placeholders
		if lesson.Type == "EMPTY" {
			continue
		}
		scheduleDay.Lessons = append(scheduleDay.Lessons, mapLesson(lesson))
	}

	return scheduleDay, nil
}

func mapLesson(dto tsuschedule.Lesson) entity.Lesson {
	return entity.Lesson{
		ID:         *dto.Id,
		Title:      *dto.Title,
		LessonType: values.LessonType(*dto.LessonType),
		Position:   dto.LessonNumber - 1,
		Audience:   mapAudience(dto.Audience),
		Groups:     mapGroups(dto.Groups),
	}
}

func mapGroups(groups *[]tsuschedule.GroupInfo) []entity.GroupInfo {
	if groups == nil {
		return nil
	}

	var result = make([]entity.GroupInfo, len(*groups))
	for i, dto := range *groups {
		result[i] = entity.GroupInfo{
			ID:   dto.Id,
			Name: dto.Name,
		}
	}
	return result
}

func mapAudience(audience *tsuschedule.AudienceInfo) entity.AudienceInfo {
	if audience == nil {
		return entity.AudienceInfo{
			ID:   "",
			Name: "Не определено",
		}
	}

	var id string
	if audience.Id == nil {
		id = ""
	}

	return entity.AudienceInfo{
		ID:   id,
		Name: audience.Name,
	}
}
