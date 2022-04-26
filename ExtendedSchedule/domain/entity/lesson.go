package entity

import (
	"tellmeac/extended-schedule/domain/values"
)

type Lesson struct {
	ID         string
	Title      string
	LessonType values.LessonType
	Position   int
	Teacher    TeacherInfo
	Audience   AudienceInfo
	Groups     []GroupInfo
}

func (lesson Lesson) Equal(other *Lesson) bool {
	return lesson.ID == other.ID &&
		lesson.Position == other.Position &&
		lesson.Teacher == other.Teacher &&
		lesson.LessonType == other.LessonType
}

type AudienceInfo struct {
	ID   string
	Name string
}

type GroupInfo struct {
	ID   string
	Name string
}

type TeacherInfo struct {
	ID   string
	Name string
}
