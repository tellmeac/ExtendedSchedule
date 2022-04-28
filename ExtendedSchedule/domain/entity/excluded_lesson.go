package entity

import (
	"github.com/google/uuid"
)

type ExcludedLesson struct {
	ID       uuid.UUID    `json:"ID"`
	LessonID string       `json:"LessonID"`
	Teacher  *TeacherInfo `json:"Teacher"`
	Position int          `json:"Position"`
	WeekDay  int          `json:"WeekDay"`
}
