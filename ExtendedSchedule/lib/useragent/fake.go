package useragent

import (
	"context"
	ua "github.com/wux1an/fake-useragent"
	"net/http"
)

// ApplyFakeUserAgent applies fake User-Agent header in request.
func ApplyFakeUserAgent(_ context.Context, request *http.Request) error {
	if request == nil {
		return nil
	}
	request.Header.Set("User-Agent", ua.RandomType(ua.Desktop))
	return nil
}
