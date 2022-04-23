package factory

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"tellmeac/extended-schedule/domain/entity"
	"testing"
	"time"
)

func TestUsersScheduleFactory_Make(t *testing.T) {
	testCases := []struct {
		Name             string
		Start            time.Time
		End              time.Time
		GroupSchedule    map[string][]entity.DaySchedule
		JoinedGroups     []entity.GroupInfo
		ExcludedLessons  []entity.ExcludedLesson
		ExtendedLessons  []entity.ExtendedLesson
		ExpectedSchedule []entity.DaySchedule
		IsExpectedError  bool
	}{
		{
			Name:  "Joined groups should be added successfully",
			Start: time.Date(2022, 4, 18, 0, 0, 0, 0, time.UTC),
			End:   time.Date(2022, 4, 18, 0, 0, 0, 0, time.UTC),
			JoinedGroups: []entity.GroupInfo{
				{
					ID: "group-1",
				},
				{
					ID: "group-2",
				},
			},
			GroupSchedule: map[string][]entity.DaySchedule{
				"group-1": {
					{
						Date: time.Date(2022, 4, 18, 0, 0, 0, 0, time.UTC),
						Sections: []entity.Section{
							{
								Position: 1,
								Lessons: []entity.Lesson{
									{
										ID:           "id-1",
										LessonType:   "PRACTICE",
										LessonNumber: 1,
									},
									{
										ID:           "id-2",
										LessonType:   "PRACTICE",
										LessonNumber: 1,
									},
								},
							},
						},
					},
				},
				"group-2": {
					{
						Date: time.Date(2022, 4, 18, 0, 0, 0, 0, time.UTC),
						Sections: []entity.Section{
							{
								Position: 1,
								Lessons: []entity.Lesson{
									{
										ID:           "id-1",
										LessonType:   "PRACTICE",
										LessonNumber: 1,
									},
								},
							},
						},
					},
				},
			},
			IsExpectedError: false,
			ExpectedSchedule: []entity.DaySchedule{
				{
					Date: time.Date(2022, 4, 18, 0, 0, 0, 0, time.UTC),
					Sections: []entity.Section{
						{
							Position: 1,
							Lessons: []entity.Lesson{
								{
									ID:           "id-1",
									LessonType:   "PRACTICE",
									LessonNumber: 1,
								},
								{
									ID:           "id-2",
									LessonType:   "PRACTICE",
									LessonNumber: 1,
								},
							},
						},
					},
				},
			},
		},
	}

	userID := uuid.New()

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			require.NotPanics(
				t,
				func() {
					scheduleProvider := new(scheduleProviderMock)
					for _, g := range testCase.JoinedGroups {
						scheduleProvider.On("GetByGroup", g.ID, testCase.Start, testCase.End).
							Return(testCase.GroupSchedule[g.ID], nil)
					}

					groupsRepository := new(joinedGroupsRepository)
					groupsRepository.On("GetByUserID", userID).Return(testCase.JoinedGroups, nil)

					excludedRepository := new(excludedLessonsRepository)
					excludedRepository.On("GetByUserID", userID).Return(testCase.ExcludedLessons, nil)

					extendedRepository := new(extendedLessonsRepository)
					extendedRepository.On("GetByUserID", userID).Return(testCase.ExtendedLessons, nil)

					factory := UsersScheduleFactory{
						scheduleProvider: scheduleProvider,
						joinedGroups:     groupsRepository,
						extendedLessons:  extendedRepository,
						excludedLessons:  excludedRepository,
					}

					result, err := factory.Make(context.Background(), userID, testCase.Start, testCase.End)
					require.Equalf(
						t,
						testCase.IsExpectedError,
						err != nil,
						fmt.Sprintf("expected error: %t", testCase.IsExpectedError),
					)

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

func (s *scheduleProviderMock) GetByGroup(_ context.Context, groupID string, start time.Time, end time.Time) ([]entity.DaySchedule, error) {
	result := s.Called(groupID, start, end)
	return result.Get(0).([]entity.DaySchedule), result.Error(1)
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

type extendedLessonsRepository struct {
	mock.Mock
}

func (e *extendedLessonsRepository) GetByUserID(_ context.Context, userID uuid.UUID) ([]entity.ExtendedLesson, error) {
	result := e.Called(userID)
	return result.Get(0).([]entity.ExtendedLesson), result.Error(1)
}

func (e *extendedLessonsRepository) Update(_ context.Context, userID uuid.UUID, desired []entity.ExtendedLesson) error {
	result := e.Called(userID, desired)
	return result.Error(0)
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
