package schedule

import (
	commonmodels "tellmeac/extended-schedule/common/models"
	"tellmeac/extended-schedule/pkg/tsuschedule"
	"time"
)

func toCommonDay(day tsuschedule.DaySchedule) (commonmodels.DaySchedule, error) {
	date, err := time.Parse(scheduleDateFormat, day.Date)
	if err != nil {
		return commonmodels.DaySchedule{}, err
	}

	scheduleDay := commonmodels.DaySchedule{
		Date:    date,
		Lessons: make([]commonmodels.LessonWithContext, 0, len(day.Lessons)),
	}

	for _, lesson := range day.Lessons {
		// ignoring empty placeholders received from provider api
		if lesson.Type == "EMPTY" {
			continue
		}
		scheduleDay.Lessons = append(scheduleDay.Lessons, toCommonLessonWithContext(lesson))
	}

	return scheduleDay, nil
}

func toCommonLessonWithContext(dto tsuschedule.Lesson) commonmodels.LessonWithContext {
	return commonmodels.LessonWithContext{
		ID:         *dto.Id,
		Title:      *dto.Title,
		LessonType: commonmodels.LessonType(*dto.LessonType),
		Position:   dto.LessonNumber - 1,
		Teacher:    toCommonTeacherInfo(dto.Professor),
		Audience:   toCommonLessonAudience(dto.Audience),
		Groups:     toCommonGroupInfo(dto.Groups),
	}
}

func toCommonTeacherInfo(teacher *tsuschedule.TeacherInfo) commonmodels.TeacherInfo {
	if teacher != nil && teacher.Id != nil && teacher.FullName != nil {
		return commonmodels.TeacherInfo{
			ID:   *teacher.Id,
			Name: *teacher.FullName,
		}
	}
	return commonmodels.TeacherInfo{}
}

func toCommonGroupInfo(groups *[]tsuschedule.GroupInfo) []commonmodels.GroupInfo {
	if groups == nil {
		return nil
	}

	var result = make([]commonmodels.GroupInfo, len(*groups))
	for i, dto := range *groups {
		result[i] = commonmodels.GroupInfo{
			ID:   dto.Id,
			Name: dto.Name,
		}
	}
	return result
}

func toCommonLessonAudience(audience *tsuschedule.AudienceInfo) commonmodels.AudienceInfo {
	if audience == nil {
		return commonmodels.AudienceInfo{}
	}

	var id string
	if audience.Id == nil {
		id = ""
	}

	return commonmodels.AudienceInfo{
		ID:   id,
		Name: audience.Name,
	}
}
