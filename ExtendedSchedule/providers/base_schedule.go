package providers

import (
	"context"
	"fmt"
	"tellmeac/extended-schedule/clients/tsuschedule"
	"time"
)

type BaseScheduleProvider struct {
	client *tsuschedule.Client
}

func (provider *BaseScheduleProvider) GetByGroup(
	ctx context.Context,
	groupID string,
	start time.Time,
	end time.Time,
) ([]tsuschedule.DaySchedule, error) {
	params := tsuschedule.GetScheduleGroupParams{
		Id:       groupID,
		DateFrom: start.Format("2006-01-02"),
		DateTo:   end.Format("2006-01-02"),
	}

	response, err := provider.client.GetScheduleGroup(ctx, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to get response from api: %w", err)
	}

	result, err := tsuschedule.ParseGetScheduleGroupResponse(response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response from api: %w", err)
	}

	if result.JSON200 == nil {
		return nil, fmt.Errorf("failed to get schedule from parsed response: %w", err)
	}

	return *result.JSON200, nil
}
