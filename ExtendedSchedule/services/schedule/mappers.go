package schedule

import (
	"tellmeac/extended-schedule/domain/lesson"
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
		Lessons: make([]schedule.Lesson, 0, len(day.Lessons)),
	}

	for _, l := range day.Lessons {
		// ignoring empty placeholders received from provider api
		if l.Type == "EMPTY" {
			continue
		}
		scheduleDay.Lessons = append(scheduleDay.Lessons, toCommonLessonWithContext(l))
	}

	return scheduleDay, nil
}

func toCommonLessonWithContext(dto tsuschedule.Lesson) schedule.Lesson {
	return schedule.Lesson{
		Lesson: lesson.Lesson{
			ID:       *dto.Id,
			Title:    *dto.Title,
			Kind:     lesson.Kind(*dto.LessonType),
			Teacher:  toCommonTeacherInfo(dto.Professor),
			Audience: toCommonLessonAudience(dto.Audience),
			Groups:   toCommonGroupInfo(dto.Groups),
		},
		Position: dto.LessonNumber - 1,
	}
}

func toCommonTeacherInfo(teacher *tsuschedule.TeacherInfo) values.Teacher {
	if teacher != nil && teacher.Id != nil && teacher.FullName != nil {
		return values.Teacher{
			ExternalID: *teacher.Id,
			Name:       *teacher.FullName,
		}
	}
	return values.Teacher{}
}

func toCommonGroupInfo(groups *[]tsuschedule.GroupInfo) []values.StudyGroup {
	if groups == nil {
		return nil
	}

	var result = make([]values.StudyGroup, len(*groups))
	for i, dto := range *groups {
		result[i] = values.StudyGroup{
			ExternalID: dto.Id,
			Name:       dto.Name,
		}
	}
	return result
}

func toCommonLessonAudience(audience *tsuschedule.AudienceInfo) values.Audience {
	if audience == nil {
		return values.Audience{}
	}

	var id string
	if audience.Id == nil {
		id = ""
	}

	return values.Audience{
		ExternalID: id,
		Name:       audience.Name,
	}
}
