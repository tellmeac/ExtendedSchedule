package bootstrap

import (
	"github.com/gin-gonic/gin"
	"tellmeac/extended-schedule/service/groups"
	"tellmeac/extended-schedule/service/schedule"
	"tellmeac/extended-schedule/service/userconfig"
	"tellmeac/extended-schedule/utils/middleware"
	"tellmeac/extended-schedule/utils/routes/health"
)

// bindRoutes binds api endpoints logically.
func bindRoutes(engine *gin.Engine, schedule *schedule.Endpoints, configs *userconfig.Endpoints, groups *groups.Endpoints) {
	health.Endpoints.Bind(engine)

	api := engine.Group("api/")

	api.Use(middleware.GoogleOAuth2())
	schedule.Bind(api)
	configs.Bind(api)
	groups.Bind(api)
}
