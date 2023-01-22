package tsuclient

import (
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"tellmeac/extended-schedule/config"
	"tellmeac/extended-schedule/pkg/utils"
)

func NewTypedClient() ClientWithResponsesInterface {
	c, err := NewClientWithResponses(config.Get().BaseScheduleUrl, func(client *Client) error {
		client.RequestEditors = append(client.RequestEditors, utils.ApplyFakeUserAgent)
		return nil
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to setup tsuclient client")
	}
	return c
}

// Module provides api client for tsuclient sc.
var Module = fx.Options(
	fx.Provide(NewTypedClient),
)
