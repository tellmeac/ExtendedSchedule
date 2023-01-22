package utils

import (
	"context"
	ua "github.com/wux1an/fake-useragent"
	"net/http"
)

// ApplyFakeUserAgent applies fake User-Agent header in http.Request.
func ApplyFakeUserAgent(_ context.Context, r *http.Request) error {
	if r == nil {
		return nil
	}
	r.Header.Set("User-Agent", ua.RandomType(ua.Desktop))
	return nil
}
