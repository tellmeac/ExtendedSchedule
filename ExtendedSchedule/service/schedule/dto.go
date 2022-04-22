package schedule

import (
	"tellmeac/extended-schedule/domain/entity"
	"time"
)

type Lesson struct {
	ID     string
	Name   string
	Type   string
	Groups []entity.GroupInfo
}

type Section struct {
	Lessons []Lesson
}

type Day struct {
	Sections []Section
}

type WeekdaySchedule struct {
	Start time.Time
	End   time.Time
	Days  []Day
}
