package tsuschedule

import (
	"go.uber.org/fx"
	"tellmeac/extended-schedule/lib/useragent"
	"tellmeac/extended-schedule/pkg/config"
)

var Module = fx.Options(fx.Provide(NewBaseScheduleClient))

// NewBaseScheduleClient attempts to make api client for tsu schedule api.
func NewBaseScheduleClient(cfg config.Config) *Client {
	options := func(client *Client) error {
		client.RequestEditors = append(client.RequestEditors, useragent.ApplyFakeUserAgent)
		return nil
	}

	client, err := NewClient(cfg.ScheduleAPI.BaseURL, options)
	if err != nil {
		panic(err)
	}

	return client
}
