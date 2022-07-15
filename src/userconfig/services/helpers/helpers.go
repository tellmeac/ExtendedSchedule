package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// DateFormat for responses and parameters.
const DateFormat = "2006-02-01"

// ExtractIntervalFromQuery tries to get valid query parameters "start" and "end" with required format.
func ExtractIntervalFromQuery(ctx *gin.Context) (start time.Time, end time.Time, err error) {
	start, err = time.Parse(DateFormat, ctx.Query("start"))
	if err != nil {
		return
	}

	end, err = time.Parse(DateFormat, ctx.Query("end"))
	if err != nil {
		return
	}

	if end.Sub(start) < 0 {
		err = errors.New("end time is less than start")
		return
	}
	return
}

// HandleBadRequest writes to response specific status code and json body.
func HandleBadRequest(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusBadRequest, gin.H{"Message": msg})
}
