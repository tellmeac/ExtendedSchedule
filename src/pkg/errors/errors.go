package errors

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SendError tries to handle error properly or as internal server error to response.
func SendError(ctx *gin.Context, err error, errorMap map[error]int) {
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
