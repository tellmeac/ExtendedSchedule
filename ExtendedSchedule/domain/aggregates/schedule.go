package aggregates

import (
	"errors"
	"tellmeac/extended-schedule/domain/entity"
	"time"
)

type DaySchedule struct {
	Date     time.Time
	Sections []Section
}

func (day *DaySchedule) ExcludeLessons(lessons []entity.ExcludedLesson) error {
	var excluded = make(map[string]*entity.ExcludedLesson, len(lessons))
	for _, lesson := range lessons {
		excluded[lesson.LessonRef.LessonID] = &lesson
	}

	var filteredLessons = make([]entity.Lesson, 0)
	for _, section := range day.Sections {
		for _, lesson := range section.Lessons {
			if e, ok := excluded[lesson.ID]; !ok || ok && !day.isExcluded(lesson, e, section.Position) {
				filteredLessons = append(filteredLessons, lesson)
			}
		}
		section.Lessons = filteredLessons
	}
	return nil
}

func (day *DaySchedule) Join(other DaySchedule) error {
	if day.Date != other.Date {
		return errors.New("expected to have equal date of day index by index")
	}

	day.Sections = JoinSections(day.Sections, other.Sections)
	return nil
}

func (day *DaySchedule) isExcluded(lesson entity.Lesson, excluded *entity.ExcludedLesson, position int) bool {
	var isIncludedWeekday = func(day int, days []int) bool {
		if days == nil {
			return true
		}

		for _, d := range days {
			if d == day {
				return true
			}
		}
		return false
	}

	if isIncludedWeekday(int(day.Date.Weekday()), excluded.ByWeekDays) && excluded.ByPosition == position {
		// assert  excluding by teacher
		if excluded.ByTeacher == nil || (excluded.ByTeacher != nil && lesson.Teacher == *excluded.ByTeacher) {
			return false
		}
	}

	return true
}

// JoinSchedules объединяет два одинаковых по размеру и дате списка расписания.
func JoinSchedules(a []DaySchedule, b []DaySchedule) ([]DaySchedule, error) {
	if a == nil {
		return b, nil
	}
	if b == nil {
		return a, nil
	}

	if len(a) != len(b) {
		return nil, errors.New("expected to have equal schedule length")
	}

	var joinedResult = a
	for i := 0; i < len(a); i++ {
		if a[i].Date != b[i].Date {
			return nil, errors.New("expected to have equal date of day index by index")
		}
		joinedResult[i].Sections = JoinSections(joinedResult[i].Sections, b[i].Sections)
	}

	return joinedResult, nil
}
