package adapters

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"tellmeac/extended-schedule/config"
	"testing"
)

func TestStudyGroupRepository_SearchGroups(t *testing.T) {
	databaseAddr := os.Getenv("TEST_DATABASE")
	if databaseAddr == "" {
		t.Skip("Database is not ready for testing")
	}

	config.Set(config.Config{
		Debug:           true,
		DatabaseAddress: databaseAddr,
	})
	client := NewEntClient()

	repository := NewStudyGroupRepository(client)

	_, err := repository.SearchGroups(context.Background(), "931901")

	assert.NoError(t, err)
}
