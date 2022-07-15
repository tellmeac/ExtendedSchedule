package schedule

import (
	"errors"
	"github.com/samber/lo"
	"github.com/tellmeac/ExtendedSchedule/userconfig/domain/userconfig"
	"github.com/tellmeac/ExtendedSchedule/userconfig/domain/values"
	"reflect"
	"time"
)

type DaySchedule struct {
	Date    time.Time `json:"date"`
	Lessons []Lesson  `json:"lessons"`
}

func (day *DaySchedule) ExcludeLessons(excluded []userconfig.ExcludeRule) error {
	var ruleMap = make(map[string][]userconfig.ExcludeRule, len(excluded))
	for _, e := range excluded {
		ruleMap[e.LessonID] = append(ruleMap[e.LessonID], e)
	}

	var filteredLessons = make([]Lesson, 0)
	for _, lesson := range day.Lessons {
		if !day.isExcluded(&lesson, ruleMap) {
			filteredLessons = append(filteredLessons, lesson)
		}
	}
	day.Lessons = filteredLessons

	return nil
}

func (day DaySchedule) isExcluded(lesson *Lesson, ruleMap map[string][]userconfig.ExcludeRule) bool {
	matchGroups := func(a []values.StudyGroup, b []values.StudyGroup) bool {
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
	var appendDistinct = func(others []*Lesson, lesson *Lesson) []*Lesson {
		for _, other := range others {
			if lesson.IsSame(other) {
				return others
			}
		}

		return append(others, lesson)
	}

	if day.Date != other.Date {
		return errors.New("expected to have equal date of day index by index")
	}

	var maxPosition = 0
	var lessonsByPosition = make(map[int][]*Lesson, 0)
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

	var joinedLessons = make([]Lesson, 0, maximumLessons)
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

// Join joins many schedule to one if they equal by length and date period.
func Join(schedule ...[]DaySchedule) ([]DaySchedule, error) {
	if len(schedule) == 0 {
		return nil, nil
	}

	isSameLength := lo.EveryBy(schedule, func(s []DaySchedule) bool {
		return len(s) == len(schedule[0])
	})
	if !isSameLength {
		return nil, errors.New("expected to have equal schedule length")
	}

	var err error
	var result = schedule[0]
	for i := 1; i < len(schedule); i++ {
		result, err = Join2(result, schedule[i])
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func Join2(a []DaySchedule, b []DaySchedule) ([]DaySchedule, error) {
	if a == nil {
		return b, nil
	}

	if b == nil {
		return a, nil
	}

	for i := 0; i < len(a); i++ {
		if a[i].Date != b[i].Date {
			return nil, errors.New("expected to have equal day dates index by index")
		}
		if err := a[i].Join(b[i]); err != nil {
			return nil, err
		}
	}

	return a, nil
}
