package tsu

import (
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"tellmeac/extended-schedule/common/utils"
	"tellmeac/extended-schedule/config"
)

func NewTypedClient() ClientWithResponsesInterface {
	c, err := NewClientWithResponses(config.Get().BaseScheduleUrl, func(client *Client) error {
		client.RequestEditors = append(client.RequestEditors, utils.ApplyFakeUserAgent)
		return nil
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to setup tsu client")
	}
	return c
}

// Module provides api client for tsu sc.
var Module = fx.Options(
	fx.Provide(NewTypedClient),
)
