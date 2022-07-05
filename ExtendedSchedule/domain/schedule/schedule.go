package schedule

import (
	"errors"
	"reflect"
	"tellmeac/extended-schedule/domain/userconfig"
	"tellmeac/extended-schedule/domain/values"
	"time"
)

type DaySchedule struct {
	Date    time.Time          `json:"date"`
	Lessons []LessonInSchedule `json:"lessons"`
}

func (day *DaySchedule) ExcludeLessons(excluded []userconfig.ExcludeRule) error {
	var ruleMap = make(map[string][]userconfig.ExcludeRule, len(excluded))
	for _, e := range excluded {
		ruleMap[e.LessonID] = append(ruleMap[e.LessonID], e)
	}

	var filteredLessons = make([]LessonInSchedule, 0)
	for _, lesson := range day.Lessons {
		if !day.isExcluded(&lesson, ruleMap) {
			filteredLessons = append(filteredLessons, lesson)
		}
	}
	day.Lessons = filteredLessons

	return nil
}

func (day DaySchedule) isExcluded(lesson *LessonInSchedule, ruleMap map[string][]userconfig.ExcludeRule) bool {
	matchGroups := func(a []values.GroupInfo, b []values.GroupInfo) bool {
		return reflect.DeepEqual(a, b)
	}

	var lessonMatchExcluded = func(excluded *userconfig.ExcludeRule) bool {
		return lesson.ID == excluded.LessonID &&
			lesson.Position == excluded.Position &&
			int(day.Date.Weekday()) == excluded.WeekDay && matchGroups(excluded.Groups, lesson.Groups)
	}

	if excludedLessons, ok := ruleMap[lesson.ID]; ok {
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
	var appendDistinct = func(others []*LessonInSchedule, lesson *LessonInSchedule) []*LessonInSchedule {
		for _, other := range others {
			if lesson.IsDuplicate(other) {
				return others
			}
		}

		return append(others, lesson)
	}

	if day.Date != other.Date {
		return errors.New("expected to have equal date of day index by index")
	}

	var maxPosition = 0
	var lessonsByPosition = make(map[int][]*LessonInSchedule, 0)
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

	var joinedLessons = make([]LessonInSchedule, 0, maximumLessons)
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

// JoinSchedules joins to schedule if they equal by length and date period.
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
