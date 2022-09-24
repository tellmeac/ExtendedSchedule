package userconfig

import (
	"go.uber.org/fx"
	"tellmeac/extended-schedule/common/utils"
	"tellmeac/extended-schedule/config"
)

// Module provides api client for user config service.
var Module = fx.Options(
	fx.Provide(func() (ClientWithResponsesInterface, error) {
		return NewClientWithResponses(config.Get().BaseScheduleUrl, func(client *Client) error {
			client.RequestEditors = append(client.RequestEditors, utils.ApplyFakeUserAgent)
			return nil
		})
	}),
)
