package schedule

import (
	"tellmeac/extended-schedule/domain/schedule"
	"tellmeac/extended-schedule/domain/values"
	"tellmeac/extended-schedule/services/tsuschedule"
	"time"
)

func toCommonDay(day tsuschedule.DaySchedule) (schedule.DaySchedule, error) {
	date, err := time.Parse(scheduleDateFormat, day.Date)
	if err != nil {
		return schedule.DaySchedule{}, err
	}

	scheduleDay := schedule.DaySchedule{
		Date:    date,
		Lessons: make([]schedule.LessonInSchedule, 0, len(day.Lessons)),
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

func toCommonLessonWithContext(dto tsuschedule.Lesson) schedule.LessonInSchedule {
	return schedule.LessonInSchedule{
		ID:         *dto.Id,
		Title:      *dto.Title,
		LessonType: schedule.LessonType(*dto.LessonType),
		Position:   dto.LessonNumber - 1,
		Teacher:    toCommonTeacherInfo(dto.Professor),
		Audience:   toCommonLessonAudience(dto.Audience),
		Groups:     toCommonGroupInfo(dto.Groups),
	}
}

func toCommonTeacherInfo(teacher *tsuschedule.TeacherInfo) values.TeacherInfo {
	if teacher != nil && teacher.Id != nil && teacher.FullName != nil {
		return values.TeacherInfo{
			ID:   *teacher.Id,
			Name: *teacher.FullName,
		}
	}
	return values.TeacherInfo{}
}

func toCommonGroupInfo(groups *[]tsuschedule.GroupInfo) []values.GroupInfo {
	if groups == nil {
		return nil
	}

	var result = make([]values.GroupInfo, len(*groups))
	for i, dto := range *groups {
		result[i] = values.GroupInfo{
			ID:   dto.Id,
			Name: dto.Name,
		}
	}
	return result
}

func toCommonLessonAudience(audience *tsuschedule.AudienceInfo) values.AudienceInfo {
	if audience == nil {
		return values.AudienceInfo{}
	}

	var id string
	if audience.Id == nil {
		id = ""
	}

	return values.AudienceInfo{
		ID:   id,
		Name: audience.Name,
	}
}
