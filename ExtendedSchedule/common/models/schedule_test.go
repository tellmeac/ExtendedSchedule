package models

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestDaySchedule_ExcludeLessons(t *testing.T) {
	scheduleDay := time.Date(2022, 4, 18, 0, 0, 0, 0, time.UTC)

	testCases := []struct {
		Name         string
		Day          DaySchedule
		ExcludeRules []ExcludeRule
		ExpectedDay  DaySchedule
	}{
		{
			Name: "ExcludeRules properly",
			Day: DaySchedule{
				Date: scheduleDay,
				Lessons: []LessonWithContext{
					{
						ID:       "abc",
						Position: 0,
						Title:    "Lesson-1",
					},
					{
						ID:       "zxc",
						Position: 1,
						Title:    "Lesson-1",
					},
				},
			},
			ExcludeRules: []ExcludeRule{
				{
					LessonID: "zxc",
					Position: 1,
					WeekDay:  int(scheduleDay.Weekday()),
				},
			},
			ExpectedDay: DaySchedule{
				Date: scheduleDay,
				Lessons: []LessonWithContext{
					{
						ID:    "abc",
						Title: "Lesson-1",
					},
				},
			},
		},
		{
			Name: "Empty excluded",
			Day: DaySchedule{
				Date: scheduleDay,
				Lessons: []LessonWithContext{
					{
						ID:    "abc",
						Title: "Lesson-1",
					},
				},
			},
			ExcludeRules: nil,
			ExpectedDay: DaySchedule{
				Date: scheduleDay,
				Lessons: []LessonWithContext{
					{
						ID:    "abc",
						Title: "Lesson-1",
					},
				},
			},
		},
		{
			Name: "ExcludeRules many",
			Day: DaySchedule{
				Date: scheduleDay,
				Lessons: []LessonWithContext{
					{
						ID:       "zxc",
						Position: 1,
						Title:    "Lesson-1",
						Teacher: TeacherInfo{
							ID:   "other-teacher",
							Name: "Teacher",
						},
					},
					{
						ID:       "zxc",
						Position: 1,
						Title:    "Lesson-1",
						Teacher: TeacherInfo{
							ID:   "teacher",
							Name: "Teacher",
						},
					},
					{
						ID:       "zxc",
						Position: 2,
						Title:    "Lesson-1",
						Teacher: TeacherInfo{
							ID:   "teacher",
							Name: "Teacher",
						},
					},
				},
			},
			ExcludeRules: []ExcludeRule{
				{
					LessonID: "zxc",
					Position: 1,
					WeekDay:  int(scheduleDay.Weekday()),
				},
			},
			ExpectedDay: DaySchedule{
				Date: scheduleDay,
				Lessons: []LessonWithContext{
					{
						ID:       "zxc",
						Position: 2,
						Title:    "Lesson-1",
						Teacher: TeacherInfo{
							ID:   "teacher",
							Name: "Teacher",
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			require.NotPanics(t, func() {
				err := tc.Day.ExcludeLessons(tc.ExcludeRules)
				require.NoError(t, err)
				require.Equal(t, tc.ExpectedDay, tc.Day)
			})
		})
	}
}

func TestDaySchedule_Join(t *testing.T) {
	scheduleDay := time.Date(2022, 4, 18, 0, 0, 0, 0, time.UTC)

	testCases := []struct {
		Name     string
		First    DaySchedule
		Second   DaySchedule
		Expected DaySchedule
	}{
		{
			Name: "Join with empty day",
			First: DaySchedule{
				Date: scheduleDay,
				Lessons: []LessonWithContext{
					{
						Title:    "Lesson",
						Position: 1,
					},
				},
			},
			Second: DaySchedule{
				Date:    scheduleDay,
				Lessons: nil,
			},
			Expected: DaySchedule{
				Date: scheduleDay,
				Lessons: []LessonWithContext{
					{
						Title:    "Lesson",
						Position: 1,
					},
				},
			},
		},
		{
			Name: "Join days with no collides",
			First: DaySchedule{
				Date: scheduleDay,
				Lessons: []LessonWithContext{
					{
						Title:    "Lesson",
						Position: 1,
					},
				},
			},
			Second: DaySchedule{
				Date: scheduleDay,
				Lessons: []LessonWithContext{
					{
						Title:    "Lesson",
						Position: 2,
					},
				},
			},
			Expected: DaySchedule{
				Date: scheduleDay,
				Lessons: []LessonWithContext{
					{
						Title:    "Lesson",
						Position: 1,
					},
					{
						Title:    "Lesson",
						Position: 2,
					},
				},
			},
		},
		{
			Name: "Join days with collides",
			First: DaySchedule{
				Date: scheduleDay,
				Lessons: []LessonWithContext{
					{
						Title:    "Lesson",
						Position: 1,
					},
				},
			},
			Second: DaySchedule{
				Date: scheduleDay,
				Lessons: []LessonWithContext{
					{
						Title:    "Lesson",
						Position: 1,
					},
				},
			},
			Expected: DaySchedule{
				Date: scheduleDay,
				Lessons: []LessonWithContext{
					{
						Title:    "Lesson",
						Position: 1,
					},
				},
			},
		},
		{
			Name: "Join days with lesson position order save",
			First: DaySchedule{
				Date: scheduleDay,
				Lessons: []LessonWithContext{
					{
						Title:    "Lesson",
						Position: 2,
					},
					{
						Title:    "Lesson",
						Position: 3,
					},
				},
			},
			Second: DaySchedule{
				Date: scheduleDay,
				Lessons: []LessonWithContext{
					{
						Title:    "Lesson",
						Position: 0,
					},
					{
						Title:    "Lesson",
						Position: 1,
					},
				},
			},
			Expected: DaySchedule{
				Date: scheduleDay,
				Lessons: []LessonWithContext{
					{
						Title:    "Lesson",
						Position: 0,
					},
					{
						Title:    "Lesson",
						Position: 1,
					},
					{
						Title:    "Lesson",
						Position: 2,
					},
					{
						Title:    "Lesson",
						Position: 3,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			require.NotPanics(t, func() {
				err := tc.First.Join(tc.Second)
				require.NoError(t, err)
				require.Equal(t, tc.Expected.Date, tc.First.Date)
				require.ElementsMatch(t, tc.Expected.Lessons, tc.First.Lessons)
				require.Equal(t, tc.Expected, tc.First)
			})
		})
	}
}
