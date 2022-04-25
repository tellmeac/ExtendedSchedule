package factory

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"tellmeac/extended-schedule/domain/aggregates"
	"tellmeac/extended-schedule/domain/entity"
	"testing"
	"time"
)

func TestUsersScheduleFactory_Make(t *testing.T) {
	userID := uuid.New()
	scheduleDay := time.Date(2022, 4, 18, 0, 0, 0, 0, time.UTC)

	testCases := []struct {
		Name             string
		Start            time.Time
		End              time.Time
		GroupSchedule    map[string][]aggregates.DaySchedule
		JoinedGroups     []entity.GroupInfo
		ExcludedLessons  []entity.ExcludedLesson
		ExpectedSchedule []aggregates.DaySchedule
		IsExpectedError  bool
	}{
		{
			Name:  "Joined groups should be added successfully",
			Start: scheduleDay,
			End:   scheduleDay,
			JoinedGroups: []entity.GroupInfo{
				{
					ID: "group-1",
				},
				{
					ID: "group-2",
				},
			},
			GroupSchedule: map[string][]aggregates.DaySchedule{
				"group-1": {
					{
						Date: scheduleDay,
						Sections: []aggregates.Section{
							{
								Position: 1,
								Lessons: []entity.Lesson{
									{
										ID:         "id-1",
										LessonType: "PRACTICE",
									},
								},
							},
						},
					},
				},
				"group-2": {
					{
						Date: scheduleDay,
						Sections: []aggregates.Section{
							{
								Position: 1,
								Lessons: []entity.Lesson{
									{
										ID:         "id-1",
										LessonType: "PRACTICE",
									},
									{
										ID:         "id-2",
										LessonType: "PRACTICE",
									},
								},
							},
						},
					},
				},
			},
			IsExpectedError: false,
			ExpectedSchedule: []aggregates.DaySchedule{
				{
					Date: scheduleDay,
					Sections: []aggregates.Section{
						{
							Position: 1,
							Lessons: []entity.Lesson{
								{
									ID:         "id-1",
									LessonType: "PRACTICE",
								},
								{
									ID:         "id-2",
									LessonType: "PRACTICE",
								},
							},
						},
					},
				},
			},
		},
		{
			Name:  "Excluded lessons should be ignored successfully",
			Start: scheduleDay,
			End:   scheduleDay,
			JoinedGroups: []entity.GroupInfo{
				{
					ID: "group-1",
				},
				{
					ID: "group-2",
				},
			},
			GroupSchedule: map[string][]aggregates.DaySchedule{
				"group-1": {
					{
						Date: scheduleDay,
						Sections: []aggregates.Section{
							{
								Position: 1,
								Lessons: []entity.Lesson{
									{
										ID:         "id-1",
										LessonType: "PRACTICE",
									},
								},
							},
						},
					},
				},
				"group-2": {
					{
						Date: scheduleDay,
						Sections: []aggregates.Section{
							{
								Position: 1,
								Lessons: []entity.Lesson{
									{
										ID:         "id-2",
										LessonType: "PRACTICE",
									},
								},
							},
						},
					},
				},
			},
			ExcludedLessons: []entity.ExcludedLesson{
				{
					UserID: userID,
					LessonRef: entity.LessonRef{
						LessonID: "id-2",
					},
					ByPosition: 1,
					ByWeekDays: nil,
					ByTeacher:  nil,
				},
			},
			IsExpectedError: false,
			ExpectedSchedule: []aggregates.DaySchedule{
				{
					Date: scheduleDay,
					Sections: []aggregates.Section{
						{
							Position: 1,
							Lessons: []entity.Lesson{
								{
									ID:         "id-1",
									LessonType: "PRACTICE",
								},
							},
						},
					},
				},
			},
		},
		{
			Name:  "Excluded lessons by teacher should be ignored successfully",
			Start: scheduleDay,
			End:   scheduleDay,
			JoinedGroups: []entity.GroupInfo{
				{
					ID: "group-1",
				},
				{
					ID: "group-2",
				},
			},
			GroupSchedule: map[string][]aggregates.DaySchedule{
				"group-1": {
					{
						Date: scheduleDay,
						Sections: []aggregates.Section{
							{
								Position: 1,
								Lessons: []entity.Lesson{
									{
										ID:         "id-1",
										LessonType: "PRACTICE",
									},
								},
							},
						},
					},
				},
				"group-2": {
					{
						Date: scheduleDay,
						Sections: []aggregates.Section{
							{
								Position: 1,
								Lessons: []entity.Lesson{
									{
										ID:         "id-2",
										LessonType: "PRACTICE",
										Teacher: entity.TeacherInfo{
											ID:   "teacher-1",
											Name: "Teacher name",
										},
									},
									{
										ID:         "id-2",
										LessonType: "PRACTICE",
										Teacher: entity.TeacherInfo{
											ID:   "teacher-2",
											Name: "",
										},
									},
								},
							},
						},
					},
				},
			},
			ExcludedLessons: []entity.ExcludedLesson{
				{
					UserID: userID,
					LessonRef: entity.LessonRef{
						LessonID: "id-2",
					},
					ByWeekDays: nil,
					ByTeacher: &entity.TeacherInfo{
						ID:   "teacher-1",
						Name: "Teacher name",
					},
				},
			},
			IsExpectedError: false,
			ExpectedSchedule: []aggregates.DaySchedule{
				{
					Date: scheduleDay,
					Sections: []aggregates.Section{
						{
							Position: 1,
							Lessons: []entity.Lesson{
								{
									ID:         "id-1",
									LessonType: "PRACTICE",
								},
								{
									ID:         "id-2",
									LessonType: "PRACTICE",
									Teacher: entity.TeacherInfo{
										ID:   "teacher-2",
										Name: "",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			Name:  "Excluded lessons by weekday should be ignored successfully",
			Start: scheduleDay,
			End:   scheduleDay,
			JoinedGroups: []entity.GroupInfo{
				{
					ID: "group-1",
				},
			},
			GroupSchedule: map[string][]aggregates.DaySchedule{
				"group-1": {
					{
						Date: scheduleDay,
						Sections: []aggregates.Section{
							{
								Position: 1,
								Lessons: []entity.Lesson{
									{
										ID:         "id-1",
										LessonType: "PRACTICE",
									},
									{
										ID:         "id-2",
										LessonType: "PRACTICE",
									},
								},
							},
							{
								Position: 2,
								Lessons: []entity.Lesson{
									{
										ID:         "id-2",
										LessonType: "PRACTICE",
									},
								},
							},
						},
					},
				},
			},
			ExcludedLessons: []entity.ExcludedLesson{
				{
					UserID: userID,
					LessonRef: entity.LessonRef{
						LessonID: "id-2",
					},
					ByPosition: 1,
					ByWeekDays: nil,
					ByTeacher:  nil,
				},
			},
			IsExpectedError: false,
			ExpectedSchedule: []aggregates.DaySchedule{
				{
					Date: scheduleDay,
					Sections: []aggregates.Section{
						{
							Position: 1,
							Lessons: []entity.Lesson{
								{
									ID:         "id-1",
									LessonType: "PRACTICE",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			require.NotPanics(
				t,
				func() {
					scheduleProvider := new(scheduleProviderMock)
					for _, g := range testCase.JoinedGroups {
						scheduleProvider.On("GetByGroupID", g.ID, testCase.Start, testCase.End).
							Return(testCase.GroupSchedule[g.ID], nil)
					}

					groupsRepository := new(joinedGroupsRepository)
					groupsRepository.On("GetByUserID", userID).Return(testCase.JoinedGroups, nil)

					excludedRepository := new(excludedLessonsRepository)
					excludedRepository.On("GetByUserID", userID).Return(testCase.ExcludedLessons, nil)

					factory := UsersScheduleFactory{
						scheduleProvider: scheduleProvider,
						joinedGroups:     groupsRepository,
						excludedLessons:  excludedRepository,
					}

					result, err := factory.Make(context.Background(), userID, testCase.Start, testCase.End)
					require.Equalf(
						t,
						testCase.IsExpectedError,
						err != nil,
						fmt.Sprintf("expected error = %t, err: %v", testCase.IsExpectedError, err),
					)
					require.ElementsMatch(t, testCase.ExpectedSchedule[0].Sections[0].Lessons, result[0].Sections[0].Lessons)
					require.ElementsMatch(t, testCase.ExpectedSchedule, result)
				},
			)
		})
	}
}

// scheduleProviderMock for mock testing.
type scheduleProviderMock struct {
	mock.Mock
}

func (s *scheduleProviderMock) GetLessonSchedule(_ context.Context, groupID string, lessonID string, start time.Time, end time.Time) ([]aggregates.DaySchedule, error) {
	result := s.Called(groupID, lessonID, start, end)
	return result.Get(0).([]aggregates.DaySchedule), result.Error(1)
}

func (s *scheduleProviderMock) GetByGroupID(_ context.Context, groupID string, start time.Time, end time.Time) ([]aggregates.DaySchedule, error) {
	result := s.Called(groupID, start, end)
	return result.Get(0).([]aggregates.DaySchedule), result.Error(1)
}

type joinedGroupsRepository struct {
	mock.Mock
}

func (j *joinedGroupsRepository) Update(_ context.Context, userID uuid.UUID, desired []entity.GroupInfo) error {
	result := j.Called(userID, desired)
	return result.Error(0)
}

func (j *joinedGroupsRepository) GetByUserID(_ context.Context, userID uuid.UUID) ([]entity.GroupInfo, error) {
	result := j.Called(userID)
	return result.Get(0).([]entity.GroupInfo), result.Error(1)
}

type excludedLessonsRepository struct {
	mock.Mock
}

func (e *excludedLessonsRepository) GetByUserID(_ context.Context, userID uuid.UUID) ([]entity.ExcludedLesson, error) {
	result := e.Called(userID)
	return result.Get(0).([]entity.ExcludedLesson), result.Error(1)
}

func (e *excludedLessonsRepository) Update(_ context.Context, userID uuid.UUID, desired []entity.ExcludedLesson) error {
	result := e.Called(userID, desired)
	return result.Error(0)
}
