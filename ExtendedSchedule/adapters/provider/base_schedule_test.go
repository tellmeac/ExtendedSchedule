package provider

import (
	"context"
	"tellmeac/extended-schedule/adapters/client/tsuschedule"
	"tellmeac/extended-schedule/config"
	"tellmeac/extended-schedule/infrastructure/log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var testBaseURL = "https://intime.tsu.ru/api/web/v1"

func TestBaseScheduleProviderIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skip integration schedule provider test")
	}

	// logging init
	log.ConfigureLogger()

	testCases := []struct {
		Name         string
		GroupID      string
		StartDate    time.Time
		EndDate      time.Time
		ExpectedDays int
	}{
		{
			Name:         "Get study week, may fail if there is no schedule at the current day",
			GroupID:      "3c9f5a5d-ffca-11eb-8169-005056bc249c",
			StartDate:    time.Date(2022, 4, 18, 0, 0, 0, 0, time.UTC),
			EndDate:      time.Date(2022, 4, 23, 0, 0, 0, 0, time.UTC),
			ExpectedDays: 6,
		},
	}

	client := tsuschedule.MakeClient(config.Config{
		ScheduleAPI: struct {
			BaseURL string
		}{
			BaseURL: testBaseURL,
		},
	})

	provider := baseScheduleProvider{
		client: client,
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			require.NotPanics(
				t,
				func() {
					result, err := provider.GetByGroupID(
						context.Background(),
						testCase.GroupID,
						testCase.StartDate,
						testCase.EndDate,
					)
					require.NoError(t, err)
					require.Equal(t, testCase.ExpectedDays, len(result))
				},
			)
		})
	}
}

func TestBaseScheduleProviderIntegration_GetLessonSchedule(t *testing.T) {
	if testing.Short() {
		t.Skip("skip integration schedule provider test")
	}

	// logging init
	log.ConfigureLogger()

	testCases := []struct {
		Name         string
		GroupID      string
		LessonID     string
		StartDate    time.Time
		EndDate      time.Time
		ExpectedDays int
	}{
		{
			Name:         "Get study week for lesson, may fail if there is no schedule at the current day",
			GroupID:      "3c9f5a5d-ffca-11eb-8169-005056bc249c",
			LessonID:     "4baeab32-7df1-496f-845f-9b733a2a3079",
			StartDate:    time.Date(2022, 4, 18, 0, 0, 0, 0, time.UTC),
			EndDate:      time.Date(2022, 4, 23, 0, 0, 0, 0, time.UTC),
			ExpectedDays: 6,
		},
	}

	client := tsuschedule.MakeClient(config.Config{
		ScheduleAPI: struct {
			BaseURL string
		}{
			BaseURL: testBaseURL,
		},
	})

	provider := baseScheduleProvider{
		client: client,
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			require.NotPanics(
				t,
				func() {
					result, err := provider.GetByGroupID(
						context.Background(),
						testCase.GroupID,
						testCase.StartDate,
						testCase.EndDate,
					)
					require.NoError(t, err)
					require.Equal(t, testCase.ExpectedDays, len(result))
				},
			)
		})
	}
}
