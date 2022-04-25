package aggregates

import (
	"github.com/stretchr/testify/require"
	"tellmeac/extended-schedule/domain/entity"
	"testing"
	"time"
)

func TestIsExcludedAssertion(t *testing.T) {
	scheduleDay := time.Date(2022, 4, 18, 0, 0, 0, 0, time.UTC)

	excluded := entity.ExcludedLesson{
		LessonRef: entity.LessonRef{
			LessonID: "A",
		},
		ByTeacher:  nil,
		ByPosition: 0,
		ByWeekDays: nil,
	}

	lessonToExclude := entity.Lesson{
		ID: "A",
	}

	pos := 1

	day := DaySchedule{
		Date: scheduleDay,
		Sections: []Section{
			{
				Position: pos,
				Lessons: []entity.Lesson{
					lessonToExclude,
					{
						ID: "B",
					},
				},
			},
		},
	}

	require.Equal(t, true, day.isExcluded(lessonToExclude, &excluded, pos))
}
