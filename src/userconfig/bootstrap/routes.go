package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/tellmeac/extended-schedule/userconfig/services/faculty"
	"github.com/tellmeac/extended-schedule/userconfig/services/lesson"
	"github.com/tellmeac/extended-schedule/userconfig/services/schedule"
	"github.com/tellmeac/extended-schedule/userconfig/services/userconfig"
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
