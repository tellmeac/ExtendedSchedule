package schedule

import (
	"tellmeac/extended-schedule/domain/entity"
)

type DaySchedule struct {
	Date    string          `json:"date"`
	Lessons []entity.Lesson `json:"lessons"`
}
