package bootstrap

import (
	"github.com/gin-gonic/gin"
	"tellmeac/extended-schedule/service/schedule"
	"tellmeac/extended-schedule/utils/routes/health"
)

func bindRoutes(s *gin.Engine, schedule *schedule.Endpoints) {
	api := s.Group("api/")
	health.Endpoints.Bind(api)
	schedule.Bind(api)
}
