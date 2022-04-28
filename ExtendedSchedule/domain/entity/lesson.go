package entity

import (
	"tellmeac/extended-schedule/domain/values"
)

type Lesson struct {
	ID         string            `json:"ID"`
	Title      string            `json:"Title"`
	LessonType values.LessonType `json:"LessonType"`
	Position   int               `json:"Position"`
	Teacher    TeacherInfo       `json:"Teacher"`
	Audience   AudienceInfo      `json:"Audience"`
	Groups     []GroupInfo       `json:"Groups"`
}

func (lesson Lesson) Equal(other *Lesson) bool {
	return lesson.ID == other.ID &&
		lesson.Position == other.Position &&
		lesson.Teacher == other.Teacher &&
		lesson.LessonType == other.LessonType
}

type AudienceInfo struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

type GroupInfo struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

type TeacherInfo struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}
