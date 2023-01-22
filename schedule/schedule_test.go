package schedule

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDay_Join(t *testing.T) {
	aDate := time.Date(2022, 7, 21, 0, 0, 0, 0, time.UTC)
	bDate := time.Date(2022, 6, 5, 0, 0, 0, 0, time.UTC)

	aID := uuid.NewString()
	bID := uuid.NewString()
	cID := uuid.NewString()

	cases := []struct {
		Name          string
		A             Day
		B             Day
		Expected      Day
		ExpectedError bool
	}{
		{
			Name: "Join with same elements",
			A: Day{
				Date: aDate,
				Lessons: []Lesson{
					{
						ID:  aID,
						Pos: 1,
					},
					{
						ID:  aID,
						Pos: 2,
					},
					{
						ID:  bID,
						Pos: 2,
					},
				},
			},
			B: Day{
				Date: aDate,
				Lessons: []Lesson{
					{
						ID:  aID,
						Pos: 1,
					},
					{
						ID:  aID,
						Pos: 2,
					},
					{
						ID:  bID,
						Pos: 2,
					},
				},
			},
			Expected: Day{
				Date: aDate,
				Lessons: []Lesson{
					{
						ID:  aID,
						Pos: 1,
					},
					{
						ID:  aID,
						Pos: 2,
					},
					{
						ID:  bID,
						Pos: 2,
					},
				},
			},
			ExpectedError: false,
		},
		{
			Name: "Join with some different elements",
			A: Day{
				Date: aDate,
				Lessons: []Lesson{
					{
						ID:  aID,
						Pos: 1,
					},
					{
						ID:  cID,
						Pos: 2,
					},
					{
						ID:  bID,
						Pos: 3,
					},
				},
			},
			B: Day{
				Date: aDate,
				Lessons: []Lesson{
					{
						ID:  bID,
						Pos: 1,
					},
					{
						ID:  aID,
						Pos: 2,
					},
					{
						ID:  bID,
						Pos: 3,
					},
				},
			},
			Expected: Day{
				Date: aDate,
				Lessons: []Lesson{
					{
						ID:  aID,
						Pos: 1,
					},
					{
						ID:  bID,
						Pos: 1,
					},
					{
						ID:  cID,
						Pos: 2,
					},
					{
						ID:  aID,
						Pos: 2,
					},
					{
						ID:  bID,
						Pos: 3,
					},
				},
			},
			ExpectedError: false,
		},
		{
			Name:          "Different dates",
			A:             Day{Date: aDate},
			B:             Day{Date: bDate},
			Expected:      Day{},
			ExpectedError: true,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			aBefore := c.A
			bBefore := c.B

			result, err := c.A.Join(c.B)

			assert.Equalf(t, c.ExpectedError, err != nil, "Check err if expected")
			assert.ElementsMatch(t, c.Expected.Lessons, result.Lessons, "Joined day's lessons should be equal to expected")
			assert.Truef(t, lo.IsSortedByKey(c.Expected.Lessons, func(l Lesson) int {
				return l.Pos
			}), "Result lessons should be ordered by position")

			// check that lesson slices is unrelated
			c.Expected.Lessons = nil

			assert.Equalf(t, aBefore, c.A, "Assert that used days didn't changed")
			assert.Equalf(t, bBefore, c.B, "Assert that used days didn't changed")
		})
	}
}
