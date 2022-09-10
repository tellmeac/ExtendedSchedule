package schedule

import (
	"time"
)

// Schedule defines model for Schedule.
type Schedule struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	// ascending ordered schedule days
	Days []Day `json:"days"`
}

// Cell defines model for Cell.
type Cell struct {
	Pos     int      `json:"pos"`
	Lessons []Lesson `json:"lessons"`
}

// Day defines model for Day.
type Day struct {
	Cells []Cell    `json:"cells"`
	Date  time.Time `json:"date"`
}

// Faculty defines model for Faculty.
type Faculty struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Lesson defines model for Lesson.
type Lesson struct {
	ID      string   `json:"id"`
	Kind    string   `json:"kind"`
	Name    string   `json:"name"`
	Teacher *Teacher `json:"teacher"`
	Groups  []string `json:"groups"`
}

// Teacher defines model for Teacher.
type Teacher struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
