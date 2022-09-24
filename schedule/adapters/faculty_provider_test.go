package adapters_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"tellmeac/extended-schedule/adapters"
	"tellmeac/extended-schedule/common/tsu"
	"testing"
)

// TestFacultyProvider_Faculties integration case.
func TestFacultyProvider_Faculties(t *testing.T) {
	client, err := tsu.NewClientWithResponses("https://intime.tsu.ru/api/web/v1")
	assert.NoError(t, err)

	provider := adapters.NewFacultyProvider(client)
	faculties, err := provider.Faculties(context.Background())

	assert.NoError(t, err)
	assert.NotEmptyf(t, faculties, "Faculty list shouldn't be empty")
}
