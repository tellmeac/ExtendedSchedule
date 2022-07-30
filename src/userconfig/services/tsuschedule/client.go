package tsuschedule

import (
	"github.com/tellmeac/extended-schedule/userconfig/config"
	"github.com/tellmeac/extended-schedule/userconfig/pkg/useragent"
)

// NewBaseScheduleClient attempts to make api client for tsu schedule api.
func NewBaseScheduleClient(cfg config.Config) *Client {
	options := func(client *Client) error {
		client.RequestEditors = append(client.RequestEditors, useragent.ApplyFakeUserAgent)
		return nil
	}

	client, err := NewClient(cfg.BaseScheduleUrl, options)
	if err != nil {
		panic(err)
	}

	return client
}
