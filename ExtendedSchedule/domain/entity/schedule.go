package entity

import "time"

const (
	EmptyLesson    = "EMPTY"
	NotEmptyLesson = "LESSON"
)

// AudienceInfo defines model for audience info.
type AudienceInfo struct {
	ID   string
	Name string
}

// GroupInfo defines model for group info.
type GroupInfo struct {
	ID   string
	Name string
}

// Lesson defines model for lesson.
type Lesson struct {
	ID           string
	Title        string
	LessonNumber int
	LessonType   string
	Audience     AudienceInfo
	Groups       []GroupInfo
}

// DaySchedule defines model for day schedule.
type DaySchedule struct {
	Date     time.Time
	Sections []Section
}

type Section struct {
	Position int
	Lessons  []Lesson
}
