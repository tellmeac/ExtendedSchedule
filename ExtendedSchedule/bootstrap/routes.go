package bootstrap

import (
	"github.com/gin-gonic/gin"
	"tellmeac/extended-schedule/services/faculty"
	"tellmeac/extended-schedule/services/lesson"
	"tellmeac/extended-schedule/services/schedule"
	"tellmeac/extended-schedule/services/userconfig"
)

// bind binds api endpoints.
func bind(
	engine *gin.Engine,
	faculty *faculty.Endpoints,
	configs *userconfig.Endpoints,
	lessons *lesson.Endpoints,
	schedule *schedule.Endpoints,
) {
	api := engine.Group("api/")

	schedule.Bind(api)
	configs.Bind(api)
	faculty.Bind(api)
	lessons.Bind(api)
}
