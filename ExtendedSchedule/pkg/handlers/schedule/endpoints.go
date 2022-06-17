package schedule

import (
	"github.com/gin-gonic/gin"
	"net/http"
	commonmodels "tellmeac/extended-schedule/common/models"
	"tellmeac/extended-schedule/lib/middleware"
	"tellmeac/extended-schedule/pkg/handlers/helpers"
	"tellmeac/extended-schedule/pkg/schedule"
)

// DaySchedule is a response object.
type DaySchedule struct {
	Date    string                           `json:"date"`
	Lessons []commonmodels.LessonWithContext `json:"lessons"`
}

// New creates new endpoints to receive schedule.
func New(manager schedule.Manager) *Endpoints {
	return &Endpoints{manager: manager}
}

type Endpoints struct {
	manager schedule.Manager
}

func (e Endpoints) Bind(router gin.IRouter) {
	router.GET("/schedule", e.GetPersonalSchedule)
	router.GET("/schedule/groups/:groupID", e.GetGroupSchedule)
}

// GetPersonalSchedule - godoc
// @Router   /api/schedule [get]
// @Summary  Get personal schedule
// @Tags     Schedule
// @Param    Authorization  header  string  true  "Authorization bearer token"
// @Param    start  query  string  true  "Start date"
// @Param    end    query  string  true  "End date"
// @Produce  application/json
// @Success  200  {array}  DaySchedule
// @Failure  401
func (e Endpoints) GetPersonalSchedule(ctx *gin.Context) {
	start, end, err := helpers.ExtractIntervalFromQuery(ctx)
	if err != nil {
		helpers.HandleBadRequest(ctx, err.Error())
		return
	}

	email, err := middleware.GetGoogleEmail(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	days, err := e.manager.GetUserScheduleByEmail(ctx, email, start, end)
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

	ctx.JSON(http.StatusOK, toResponse(days))
}

// GetGroupSchedule - godoc
// @Router   /api/schedule/groups/{groupID} [get]
// @Summary  Get group schedule
// @Tags     Schedule
// @Param    groupID  path  string  true  "group ID"
// @Param    start  query  string  true  "Start date"
// @Param    end    query  string  true  "End date"
// @Produce  application/json
// @Success  200  {array}  DaySchedule
// @Failure  401
func (e Endpoints) GetGroupSchedule(ctx *gin.Context) {
	groupID := ctx.Param("groupID")
	if groupID == "" {
		helpers.HandleBadRequest(ctx, "groupID is empty")
		return
	}

	start, end, err := helpers.ExtractIntervalFromQuery(ctx)
	if err != nil {
		helpers.HandleBadRequest(ctx, err.Error())
		return
	}

	days, err := e.manager.GetByGroup(ctx, groupID, start, end)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"Message": err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, toResponse(days))
}

func toResponse(schedule []commonmodels.DaySchedule) []DaySchedule {
	var result = make([]DaySchedule, len(schedule))
	for i, d := range schedule {
		result[i] = DaySchedule{
			Date:    d.Date.Format(helpers.DateFormat),
			Lessons: d.Lessons,
		}
	}
	return result
}
