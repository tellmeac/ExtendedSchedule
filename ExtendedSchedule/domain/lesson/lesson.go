package lesson

import "tellmeac/extended-schedule/domain/values"

// Lesson represents lesson.
type Lesson struct {
	ID       string              `json:"id"`
	Title    string              `json:"title"`
	Kind     Kind                `json:"kind"`
	Audience values.Audience     `json:"audience"`
	Teacher  values.Teacher      `json:"teacher"`
	Groups   []values.StudyGroup `json:"groups"`
}
