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

func (s Schedule) Join(other Schedule) (Schedule, error) {
	if s.StartDate != other.StartDate || s.EndDate != other.EndDate {
		return Schedule{}, errors.New("schedule has different periods")
	}

	var err error
	for i, d := range s.Days {
		s.Days[i], err = d.Join(other.Days[i])
		if err != nil {
			return Schedule{}, fmt.Errorf("failed to join day: %w", err)
		}
	}

	return s, nil
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
	// join all lessons
	d.Lessons = append(d.Lessons, other.Lessons...)

	// remove duplicates
	d.Lessons = lo.UniqBy(d.Lessons, func(l Lesson) string {
		return fmt.Sprintf("%s-%d", l.ID, l.Pos)
	})

	// order by position
	slices.SortFunc(d.Lessons, func(a, b Lesson) bool {
		return a.Pos < b.Pos
	})

	return d, nil
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
