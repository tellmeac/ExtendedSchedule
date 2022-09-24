package schedule

import (
	"errors"
	"fmt"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
	"time"
)

// Schedule defines model for Schedule.
type Schedule struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Days      []Day     `json:"days"`
}

// Day defines model for Day.
type Day struct {
	Date    time.Time `json:"date"`
	Lessons []Lesson  `json:"lessons"`
}

func (d Day) Join(other Day) (Day, error) {
	if d.Date != other.Date {
		return Day{}, errors.New("days have different date value")
	}

	// init result
	maxLessons := len(d.Lessons) + len(other.Lessons)
	result := Day{
		Date:    d.Date,
		Lessons: make([]Lesson, 0, maxLessons),
	}
	result.Lessons = append([]Lesson{}, d.Lessons...)

	// join all lessons
	result.Lessons = append(result.Lessons, other.Lessons...)

	// remove duplicates
	result.Lessons = lo.UniqBy(result.Lessons, func(l Lesson) string {
		return fmt.Sprintf("%s-%d", l.ID, l.Pos)
	})

	// order by position
	slices.SortFunc(result.Lessons, func(a, b Lesson) bool {
		return a.Pos < b.Pos
	})

	return result, nil
}

// Faculty defines model for Faculty.
type Faculty struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Lesson defines model for Lesson.
type Lesson struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Kind    string   `json:"kind"`
	Pos     int      `json:"pos"`
	Teacher *Teacher `json:"teacher"`
	Groups  []string `json:"groups"`
}

// Teacher defines model for Teacher.
type Teacher struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
