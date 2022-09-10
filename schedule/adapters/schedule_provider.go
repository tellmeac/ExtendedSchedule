package adapters

import (
	"context"
	"fmt"
	"time"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/tellmeac/ext-schedule/schedule/common/tsu"
	"github.com/tellmeac/ext-schedule/schedule/schedule"
)

func NewScheduleProvider(client tsu.ClientWithResponsesInterface) ScheduleProvider {
	return ScheduleProvider{client: client}
}

type ScheduleProvider struct {
	client tsu.ClientWithResponsesInterface
}

func (sp ScheduleProvider) GetByGroup(ctx context.Context, id string, from, to time.Time) (schedule.Schedule, error) {
	resp, err := sp.client.GetScheduleGroupWithResponse(ctx, &tsu.GetScheduleGroupParams{
		ID:       id,
		DateFrom: openapi_types.Date{Time: from},
		DateTo:   openapi_types.Date{Time: to},
	})
	switch {
	case err != nil:
		return schedule.Schedule{}, fmt.Errorf("failed to get response: %w", err)
	case resp.HTTPResponse.StatusCode != 200:
		return schedule.Schedule{}, fmt.Errorf("failed with status code = %d", resp.HTTPResponse.StatusCode)
	}

	return fromScheduleDto(resp.JSON200)
}

func fromScheduleDto(dto *[]tsu.DaySchedule) (schedule.Schedule, error) {
	return schedule.Schedule{}, nil
}
