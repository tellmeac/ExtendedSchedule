package errors

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// aliases
var (
	Is  = errors.Is
	As  = errors.As
	New = errors.New
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
	if err == nil {
		panic(errors.New("can't handle nil error"))
	}

	if statusCode, ok := errorMap[err]; ok {
		ctx.JSON(statusCode, gin.H{
			"Error": err.Error(),
		})
		return
	}

	_ = ctx.AbortWithError(http.StatusInternalServerError, err)
}
