package tsuschedule

import (
	"tellmeac/extended-schedule/config"
	"tellmeac/extended-schedule/utils/useragent"
)

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
