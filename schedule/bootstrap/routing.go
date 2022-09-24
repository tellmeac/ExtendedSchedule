package bootstrap

import (
	"github.com/gin-gonic/gin"
	"tellmeac/extended-schedule/ports/schedule"
)

func setupRouting(router *gin.Engine, scheduleHandler *schedule.ServerHandler) {
	schedule.RegisterHandlersWithOptions(router, scheduleHandler, schedule.GinServerOptions{
		BaseURL:     "api/v1",
		Middlewares: nil,
	})
}
