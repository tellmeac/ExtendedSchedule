package adapters

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"tellmeac/extended-schedule/common/userconfig"
)

func NewConfigProvider(client userconfig.ClientWithResponsesInterface) ConfigProvider {
	return ConfigProvider{client: client}
}

type ConfigProvider struct {
	client userconfig.ClientWithResponsesInterface
}

func (cp ConfigProvider) Get(ctx context.Context, userID uuid.UUID) (userconfig.UserConfig, error) {
	resp, err := cp.client.GetUsersIdConfigWithResponse(ctx, userID)
	switch {
	case err != nil:
		return userconfig.UserConfig{}, fmt.Errorf("failed to get response: %w", err)
	case resp.HTTPResponse.StatusCode != 200:
		return userconfig.UserConfig{}, fmt.Errorf("failed with status code = %d", resp.HTTPResponse.StatusCode)
	}

	return *resp.JSON200, nil
}
