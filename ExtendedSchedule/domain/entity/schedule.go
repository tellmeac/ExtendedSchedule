package entity

// AudienceInfo defines model for audience info.
type AudienceInfo struct {
	Id   *string
	Name string
}

// GroupInfo defines model for group info.
type GroupInfo struct {
	Id   string
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
	Date    *string
	Lessons *[]Lesson
}
