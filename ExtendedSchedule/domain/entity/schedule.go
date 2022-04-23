package entity

import "time"

// AudienceInfo defines model for audience info.
type AudienceInfo struct {
	ID   *string
	Name string
}

// GroupInfo defines model for group info.
type GroupInfo struct {
	ID   string
	Name string
}

// Lesson defines model for lesson.
type Lesson struct {
	Audience     *AudienceInfo
	Groups       *[]GroupInfo
	LessonNumber int
	LessonType   *string
	Title        *string
	Type         string
}

// DaySchedule defines model for day schedule.
type DaySchedule struct {
	Date    time.Time
	Lessons []Lesson
}
