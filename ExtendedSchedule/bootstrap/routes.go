package bootstrap

import (
	"github.com/gin-gonic/gin"
	"tellmeac/extended-schedule/pkg/handlers/faculty"
	"tellmeac/extended-schedule/pkg/handlers/lesson"
	"tellmeac/extended-schedule/pkg/handlers/schedule"
	"tellmeac/extended-schedule/pkg/handlers/userconfig"
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
