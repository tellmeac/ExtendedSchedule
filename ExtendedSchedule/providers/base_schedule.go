package providers

import (
	"context"
	"fmt"
	"tellmeac/extended-schedule/clients/tsuschedule"
	"tellmeac/extended-schedule/domain"
	"tellmeac/extended-schedule/domain/entity"
	"time"
)

type BaseScheduleProvider struct {
	client *tsuschedule.Client
}

func (provider *BaseScheduleProvider) GetByGroup(
	ctx context.Context,
	groupID string,
	start time.Time,
	end time.Time,
) ([]entity.DaySchedule, error) {
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

	var result = make([]entity.DaySchedule, len(*scheduleDto.JSON200))
	for i, day := range *scheduleDto.JSON200 {
		result[i], err = mapDaySchedule(day)
		if err != nil {
			return nil, fmt.Errorf("failed to map response data properly: %w", err)
		}
	}

	return result, nil
}

func mapDaySchedule(day tsuschedule.DaySchedule) (entity.DaySchedule, error) {
	date, err := time.Parse("2006-01-02", day.Date)
	if err != nil {
		return entity.DaySchedule{}, err
	}

	return entity.DaySchedule{
		Date:     date,
		Sections: makeSections(day.Lessons),
	}, nil
}

func makeSections(lessons []tsuschedule.Lesson) []entity.Section {
	var sections = make([]entity.Section, domain.TotalSections)

	for _, lesson := range lessons {
		if lesson.Type == entity.EmptyLesson {
			continue
		}

		sections[lesson.LessonNumber].Position = lesson.LessonNumber
		sections[lesson.LessonNumber].Lessons = append(sections[lesson.LessonNumber].Lessons, mapLesson(lesson))
	}

	return sections
}

func mapLesson(dto tsuschedule.Lesson) entity.Lesson {
	return entity.Lesson{
		Audience:     mapAudience(dto.Audience),
		Groups:       mapGroups(dto.Groups),
		LessonNumber: dto.LessonNumber,
		LessonType:   dto.LessonType,
		Title:        dto.Title,
		Type:         dto.Type,
	}
}

func mapGroups(groups *[]tsuschedule.GroupInfo) *[]entity.GroupInfo {
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
	return &result
}

func mapAudience(audience *tsuschedule.AudienceInfo) *entity.AudienceInfo {
	if audience == nil {
		return nil
	}

	return &entity.AudienceInfo{
		ID:   audience.Id,
		Name: audience.Name,
	}
}
