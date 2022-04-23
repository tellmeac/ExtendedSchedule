package tsuschedule

import (
	"tellmeac/extended-schedule/config"
	"tellmeac/extended-schedule/infrastructure/log"
	"tellmeac/extended-schedule/utils/useragent"
)

// MakeClient attempts to make api client to tsu schedule service or panics.
func MakeClient(cfg config.Config) *Client {
	options := func(client *Client) error {
		client.RequestEditors = append(client.RequestEditors, useragent.UseFakeUserAgent)
		return nil
	}

	client, err := NewClient(cfg.ScheduleAPI.BaseURL, options)
	if err != nil {
		log.Sugared.Error("Failed to init api client: %w", err)
	}

	return client
}
