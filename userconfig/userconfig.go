package userconfig

import "github.com/google/uuid"

func New(email string) UserConfig {
	return UserConfig{
		ID:    uuid.New(),
		Email: email,
	}
}

// UserConfig defines model for user sc config.
type UserConfig struct {
	ID             uuid.UUID       `json:"id"`
	Email          string          `json:"email"`
	Base           interface{}     `json:"base"`
	ExtendedGroups []ExtendedGroup `json:"extendedGroups"`
	ExcludeRules   []ExcludeRule   `json:"excludeRules"`
}

// ExtendedGroup defines rules for extended scheduling.
type ExtendedGroup struct {
	ID      string   `json:"id"`
	Lessons []Lesson `json:"lessons"`
}

type ExcludeRule struct {
	LessonID string `json:"lessonId"`
	Pos      *int   `json:"pos"`
}

// Lesson defines model for Lesson.
type Lesson struct {
	ID   string `json:"id"`
	Kind string `json:"kind"`
	Name string `json:"name"`
}

// StudyGroup represents group for base definition.
type StudyGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Teacher is a base definition for config.
type Teacher struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
