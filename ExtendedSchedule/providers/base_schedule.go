package providers

import (
	"context"
	"fmt"
	"tellmeac/extended-schedule/clients/tsuschedule"
	"tellmeac/extended-schedule/domain"
	"tellmeac/extended-schedule/domain/aggregates"
	"tellmeac/extended-schedule/domain/entity"
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

func filterByLessonID(schedule aggregates.DaySchedule, lessonID string) aggregates.DaySchedule {
	var filteredLessons = make([]entity.Lesson, 0)
	for _, section := range schedule.Sections {
		for _, lesson := range section.Lessons {
			if lesson.ID == lessonID {
				filteredLessons = append(filteredLessons, lesson)
			}
		}
		section.Lessons = filteredLessons
	}
	return schedule
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

	return aggregates.DaySchedule{
		Date:     date,
		Sections: makeSections(day.Lessons),
	}, nil
}

func makeSections(lessons []tsuschedule.Lesson) []aggregates.Section {
	var sections = make([]aggregates.Section, domain.TotalSections)

	for _, lesson := range lessons {
		// ignoring empty placeholders
		if lesson.Type == entity.EmptyLesson {
			continue
		}

		sections[lesson.LessonNumber-1].Position = lesson.LessonNumber
		sections[lesson.LessonNumber-1].Lessons = append(sections[lesson.LessonNumber].Lessons, mapLesson(lesson))
	}

	return sections
}

func mapLesson(dto tsuschedule.Lesson) entity.Lesson {
	return entity.Lesson{
		ID:           *dto.Id,
		Title:        *dto.Title,
		LessonType:   *dto.LessonType,
		LessonNumber: dto.LessonNumber,
		Audience:     mapAudience(dto.Audience),
		Groups:       mapGroups(dto.Groups),
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
