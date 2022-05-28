package schedule

import (
	"errors"
	"net/http"
	"tellmeac/extended-schedule/domain"
	"tellmeac/extended-schedule/domain/aggregate"
	"tellmeac/extended-schedule/utils/middleware"
	"tellmeac/extended-schedule/utils/shortcuts"
	"time"

	"github.com/gin-gonic/gin"
)

func NewEndpoints(service IService) *Endpoints {
	return &Endpoints{service: service}
}

type Endpoints struct {
	service IService
}

func (e Endpoints) Bind(router gin.IRouter) {
	router.GET("/schedule/personal", e.GetPersonalSchedule)
	router.GET("/schedule/groups/:groupID", e.GetGroupSchedule)
	router.GET("/schedule/groups/:groupID/lessons/:lessonID", e.GetLessonSchedule)
}

// GetPersonalSchedule - godoc
// @Router   /api/schedule/personal [get]
// @Summary  Get personal schedule
// @Tags     Schedule
// @Param    startDate  query  string  true  "Start date"
// @Param    endDate    query  string  true  "End date"
// @Produce  application/json
// @Success  200  {array}  DaySchedule
// @Failure  401
// @Failure  500
func (e Endpoints) GetPersonalSchedule(ctx *gin.Context) {
	start, end, err := e.ExtractScheduleQuery(ctx)
	if err != nil {
		shortcuts.HandleBadRequest(ctx, err.Error())
		return
	}

	userIdentifier, err := middleware.GetGoogleEmail(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	days, err := e.service.GetPersonal(ctx, userIdentifier, start, end)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error":   err.Error(),
				"context": "Failed to build personal schedule days: %s",
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, mapSchedule(days))
}

func (e Endpoints) GetGroupSchedule(ctx *gin.Context) {
	groupID := ctx.Param("groupID")
	if groupID == "" {
		shortcuts.HandleBadRequest(ctx, "groupID is an empty string, must be identifier")
		return
	}

	start, end, err := e.ExtractScheduleQuery(ctx)
	if err != nil {
		shortcuts.HandleBadRequest(ctx, err.Error())
		return
	}

	days, err := e.service.GetByGroup(ctx, groupID, start, end)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error":   err.Error(),
				"context": "Failed to build schedule days for group",
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, mapSchedule(days))
}

func (e Endpoints) GetLessonSchedule(ctx *gin.Context) {
	groupID := ctx.Param("groupID")
	if groupID == "" {
		shortcuts.HandleBadRequest(ctx, "groupID is an empty string, must be identifier")
		return
	}

	lessonID := ctx.Param("lessonID")
	if lessonID == "" {
		shortcuts.HandleBadRequest(ctx, "lessonID is an empty string, must be identifier")
		return
	}

	start, end, err := e.ExtractScheduleQuery(ctx)
	if err != nil {
		shortcuts.HandleBadRequest(ctx, err.Error())
		return
	}

	days, err := e.service.GetByLesson(ctx, groupID, lessonID, start, end)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error":   err.Error(),
				"context": "Failed to build schedule days for group",
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, mapSchedule(days))
}

func (e Endpoints) ExtractScheduleQuery(ctx *gin.Context) (start time.Time, end time.Time, err error) {
	start, err = time.Parse(domain.ScheduleDateFormat, ctx.Query("start"))
	if err != nil {
		return
	}

	end, err = time.Parse(domain.ScheduleDateFormat, ctx.Query("end"))
	if err != nil {
		return
	}

	if end.Sub(start) < 0 {
		err = errors.New("end is less then start")
		return
	}
	return
}

func mapSchedule(schedule []aggregate.DaySchedule) []DaySchedule {
	var result = make([]DaySchedule, len(schedule))
	for i, d := range schedule {
		result[i] = DaySchedule{
			Date:    d.Date.Format(domain.ScheduleDateFormat),
			Lessons: d.Lessons,
		}
	}
	return result
}
