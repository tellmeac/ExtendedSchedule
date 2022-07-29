package schedule

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tellmeac/extended-schedule/pkg/middleware"
	"github.com/tellmeac/extended-schedule/userconfig/services/helpers"
)

type QueryPeriod struct {
	Start time.Time `form:"start" binding:"required" time_format:"2006-01-02"`
	End   time.Time `form:"end" binding:"required" time_format:"2006-01-02"`
}

// NewEndpoints creates new endpoints to receive schedule.
func NewEndpoints(manager Service) *Endpoints {
	return &Endpoints{service: manager}
}

type Endpoints struct {
	service Service
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
	var period QueryPeriod
	if err := ctx.ShouldBindQuery(&period); err != nil {
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

	days, err := e.service.GetUserScheduleByEmail(ctx, email, period.Start, period.End)
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

	ctx.JSON(http.StatusOK, days)
}

// GetGroupSchedule - godoc
// @Router   /api/schedule/groups/{groupID} [get]
// @Summary  Get group schedule
// @Tags     Schedule
// @Param    groupID  path  string  true  "group external id"
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

	var period QueryPeriod
	if err := ctx.ShouldBindQuery(&period); err != nil {
		helpers.HandleBadRequest(ctx, err.Error())
		return
	}

	days, err := e.service.GetByGroup(ctx, groupID, period.Start, period.End)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"Message": err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, days)
}
