package useragent

import (
	"context"
	ua "github.com/wux1an/fake-useragent"
	"net/http"
)

// UseFakeUserAgent replaces user agent header value as something fake.
func UseFakeUserAgent(_ context.Context, request *http.Request) error {
	if request == nil {
		return nil
	}

	request.Header.Set("User-Agent", string(ua.RandomType(ua.Desktop)))

	return nil
}
