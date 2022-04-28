package bootstrap

import (
	"github.com/gin-gonic/gin"
	"tellmeac/extended-schedule/service/schedule"
	"tellmeac/extended-schedule/service/userconfig"
	"tellmeac/extended-schedule/utils/routes/health"
)

// bindRoutes binds api endpoints logically.
func bindRoutes(s *gin.Engine, schedule *schedule.Endpoints, configs *userconfig.Endpoints) {
	api := s.Group("api/")
	health.Endpoints.Bind(api)
	schedule.Bind(api)
	configs.Bind(api)
}
