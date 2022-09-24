package schedule

import (
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
	Lessons []Lesson  `json:"lessons"`
	Date    time.Time `json:"date"`
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
