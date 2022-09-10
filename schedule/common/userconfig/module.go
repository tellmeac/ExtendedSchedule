package userconfig

import (
	"github.com/tellmeac/ext-schedule/schedule/common/utils"
	"github.com/tellmeac/ext-schedule/schedule/config"
	"go.uber.org/fx"
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
