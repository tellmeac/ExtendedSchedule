package ports

import (
	"github.com/gin-gonic/gin"
	"tellmeac/extended-schedule/ports/schedule"
)

func ApplyRouting(router *gin.Engine, scheduleHandler *schedule.ServerHandler) {
	schedule.RegisterHandlersWithOptions(router, scheduleHandler, schedule.GinServerOptions{
		BaseURL:     "api/v1",
		Middlewares: nil, // TODO: authorization, tracing, etc
	})
}
