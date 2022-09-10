package adapters_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tellmeac/ext-schedule/schedule/adapters"
	"github.com/tellmeac/ext-schedule/schedule/common/tsu"
	"github.com/tellmeac/ext-schedule/schedule/common/utils"
)

// TestScheduleProvider_GetByGroup integration case.
func TestScheduleProvider_GetByGroup(t *testing.T) {
	client, err := tsu.NewClientWithResponses("https://intime.tsu.ru/api/web/v1")
	assert.NoError(t, err)

	id := "3c9f5a5d-ffca-11eb-8169-005056bc249c" // 931901
	from, to := utils.GetWeek(time.Now())

	provider := adapters.NewScheduleProvider(client)
	schedule, err := provider.GetByGroup(context.Background(), id, from, to)

	assert.NoError(t, err)
	assert.NotEmpty(t, schedule)
}
