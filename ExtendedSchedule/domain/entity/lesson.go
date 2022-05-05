package entity

import (
	"tellmeac/extended-schedule/domain/values"
)

type Lesson struct {
	ID         string            `json:"id"`
	Title      string            `json:"title"`
	LessonType values.LessonType `json:"lessonType"`
	Position   int               `json:"position"`
	Teacher    TeacherInfo       `json:"teacher"`
	Audience   AudienceInfo      `json:"audience"`
	Groups     []GroupInfo       `json:"groups"`
}

func (lesson Lesson) Equal(other *Lesson) bool {
	return lesson.ID == other.ID &&
		lesson.Position == other.Position &&
		lesson.Teacher == other.Teacher &&
		lesson.LessonType == other.LessonType
}

type AudienceInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GroupInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TeacherInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
