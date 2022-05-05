package schedule

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"tellmeac/extended-schedule/domain"
	"tellmeac/extended-schedule/domain/aggregates"
	"tellmeac/extended-schedule/utils/shortcuts"
	"time"
)

func NewEndpoints(service IService) *Endpoints {
	return &Endpoints{service: service}
}

type Endpoints struct {
	service IService
}

func (e Endpoints) Bind(router gin.IRouter) {
	router.GET("/schedule/personal/:userID", e.GetPersonalSchedule)
	router.GET("/schedule/groups/:groupID", e.GetGroupSchedule)
	router.GET("/schedule/groups/:groupID/lessons/:lessonID", e.GetLessonSchedule)
}

func (e Endpoints) GetPersonalSchedule(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("userID"))
	if err != nil {
		shortcuts.HandleBadRequest(ctx, err.Error())
		return
	}

	start, end, err := e.ExtractScheduleQuery(ctx)
	if err != nil {
		shortcuts.HandleBadRequest(ctx, err.Error())
		return
	}

	days, err := e.service.GetPersonal(ctx, userID, start, end)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"error":   err.Error(),
				"context": "Failed to build personal schedule days",
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

func mapSchedule(schedule []aggregates.DaySchedule) []DaySchedule {
	var result = make([]DaySchedule, len(schedule))
	for i, d := range schedule {
		result[i] = DaySchedule{
			Date:    d.Date.Format("2006-01-02"),
			Lessons: d.Lessons,
		}
	}
	return result
}
