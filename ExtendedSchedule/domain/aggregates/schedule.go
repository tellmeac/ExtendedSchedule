package aggregates

import (
	"errors"
	"tellmeac/extended-schedule/domain/entity"
	"time"
)

type DaySchedule struct {
	Date    time.Time       `json:"date"`
	Lessons []entity.Lesson `json:"lessons"`
}

func (day *DaySchedule) ExcludeLessons(excluded []entity.ExcludedLesson) error {
	var excludedMap = make(map[string][]entity.ExcludedLesson, len(excluded))
	for _, e := range excluded {
		excludedMap[e.LessonID] = append(excludedMap[e.LessonID], e)
	}

	var filteredLessons = make([]entity.Lesson, 0)
	for _, lesson := range day.Lessons {
		if !day.isExcluded(&lesson, excludedMap) {
			filteredLessons = append(filteredLessons, lesson)
		}
	}
	day.Lessons = filteredLessons

	return nil
}

func (day DaySchedule) isExcluded(lesson *entity.Lesson, excludeMap map[string][]entity.ExcludedLesson) bool {
	var lessonMatchExcluded = func(excluded *entity.ExcludedLesson) bool {
		return lesson.ID == excluded.LessonID &&
			lesson.Position == excluded.Position &&
			int(day.Date.Weekday()) == excluded.WeekDay &&
			(excluded.Teacher == nil || (excluded.Teacher != nil && *excluded.Teacher == lesson.Teacher))
	}

	if excludedLessons, ok := excludeMap[lesson.ID]; ok {
		for _, excluded := range excludedLessons {
			if lessonMatchExcluded(&excluded) {
				return true
			}
		}
		return false
	}
	return false
}

// Join добавляет расписание занятий из переданного дня.
func (day *DaySchedule) Join(other DaySchedule) error {
	var appendDistinct = func(others []*entity.Lesson, lesson *entity.Lesson) []*entity.Lesson {
		for _, other := range others {
			if lesson.Equal(other) {
				return others
			}
		}

		return append(others, lesson)
	}

	if day.Date != other.Date {
		return errors.New("expected to have equal date of day index by index")
	}

	var maxPosition = 0
	var lessonsByPosition = make(map[int][]*entity.Lesson, 0)
	for i, lesson := range day.Lessons {
		lessonsByPosition[lesson.Position] = append(lessonsByPosition[lesson.Position], &day.Lessons[i])
		if maxPosition < lesson.Position {
			maxPosition = lesson.Position
		}
	}
	for i, lesson := range other.Lessons {
		lessonsByPosition[lesson.Position] = appendDistinct(lessonsByPosition[lesson.Position], &other.Lessons[i])
		if maxPosition < lesson.Position {
			maxPosition = lesson.Position
		}
	}

	// get maximum lessons for starting allocation size
	maximumLessons := len(day.Lessons)
	if len(other.Lessons) > maximumLessons {
		maximumLessons = len(other.Lessons)
	}

	var joinedLessons = make([]entity.Lesson, 0, maximumLessons)
	for i := 0; i <= maxPosition; i++ {
		if lessons, ok := lessonsByPosition[i]; ok {
			for _, lesson := range lessons {
				joinedLessons = append(joinedLessons, *lesson)
			}
		}
	}

	day.Lessons = joinedLessons
	return nil
}

// JoinSchedules объединяет два расписания одинаковых по размеру и дате.
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
		if err := joinedResult[i].Join(b[i]); err != nil {
			return nil, err
		}
	}

	return joinedResult, nil
}
