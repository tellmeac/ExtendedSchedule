package errors

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ErrNotFound = errors.New("not found")
)

var (
	Is  = errors.Is
	As  = errors.As
	New = errors.New
)

var errMap = map[error]int{
	ErrNotFound: http.StatusNotFound,
}

// SendError tries to handle error properly or as internal server error to response.
func SendError(ctx *gin.Context, err error) {
	if err == nil {
		panic(errors.New("can't handle nil error"))
	}

	if statusCode, ok := errMap[err]; ok {
		ctx.JSON(statusCode, gin.H{
			"Error": err.Error(),
		})
		return
	}

	_ = ctx.AbortWithError(http.StatusInternalServerError, err)
}
