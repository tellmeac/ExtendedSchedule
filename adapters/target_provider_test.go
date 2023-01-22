package adapters_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"tellmeac/extended-schedule/adapters"
	"tellmeac/extended-schedule/pkg/tsuclient"
	"testing"
)

// TestFacultyProvider_Faculties integration case.
func TestFacultyProvider_Faculties(t *testing.T) {
	client, err := tsuclient.NewClientWithResponses("https://intime.tsu.ru/api/web/v1")
	assert.NoError(t, err)

	provider := adapters.NewTargetProvider(client)
	faculties, err := provider.Faculties(context.Background())

	assert.NoError(t, err)
	assert.NotEmptyf(t, faculties, "List of faculties shouldn't be empty")
}

// TestFacultyProvider_Teachers integration case.
func TestFacultyProvider_Teachers(t *testing.T) {
	client, err := tsuclient.NewClientWithResponses("https://intime.tsu.ru/api/web/v1")
	assert.NoError(t, err)

	provider := adapters.NewTargetProvider(client)
	teachers, err := provider.Teachers(context.Background())

	assert.NoError(t, err)
	assert.NotEmptyf(t, teachers, "List of teachers shouldn't be empty")
}
