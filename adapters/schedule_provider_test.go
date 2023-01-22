package adapters_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"tellmeac/extended-schedule/adapters"
	"tellmeac/extended-schedule/pkg/tsuclient"
	"tellmeac/extended-schedule/pkg/utils"
)

// TestScheduleProvider_GetByGroup integration case.
func TestScheduleProvider_GetByGroup(t *testing.T) {
	client, err := tsuclient.NewClientWithResponses("https://intime.tsu.ru/api/web/v1")
	assert.NoError(t, err)

	id := "3c9f5a5d-ffca-11eb-8169-005056bc249c" // 931901
	from, to := utils.GetWeek(time.Now())

	provider := adapters.NewScheduleProvider(client)
	schedule, err := provider.GetByGroup(context.Background(), id, from, to)

	assert.NoError(t, err)
	assert.NotEmpty(t, schedule)
}
