package errors

import (
	"github.com/gin-gonic/gin"
	"github.com/tellmeac/ExtendedSchedule/pkg/errors"
	"net/http"
)

var (
	// ErrNotFound represents object not found error.
	ErrNotFound = errors.New("object was not found")
	// ErrUnauthorized represents unauthorized access error.
	ErrUnauthorized = errors.New("unauthorized")
)

var (
	errorMap = map[error]int{
		ErrNotFound:     http.StatusNotFound,
		ErrUnauthorized: http.StatusUnauthorized,
	}
)

// SendError tries to handle err properly or as internal.
func SendError(ctx *gin.Context, err error) {
	errors.SendError(ctx, err, errorMap)
}
