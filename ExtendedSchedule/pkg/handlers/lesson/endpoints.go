package lesson

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tellmeac/extended-schedule/pkg/handlers/helpers"
	"tellmeac/extended-schedule/pkg/lesson"
)

// New creates new endpoints for lessons.
func New(manager lesson.Manager) *Endpoints {
	return &Endpoints{manager: manager}
}

type Endpoints struct {
	manager lesson.Manager
}

func (e Endpoints) Bind(router gin.IRouter) {
	router.GET("/lessons", e.GetLessonList)
}

// GetLessonList - godoc
// @Router   /api/lessons [get]
// @Summary  Get group's lesson list
// @Tags     Lessons
// @Produce  application/json
// @Success  200  {array}  models.Lesson
// @Failure  404
func (e Endpoints) GetLessonList(ctx *gin.Context) {
	groupID := ctx.Query("groupId")
	if groupID == "" {
		helpers.HandleBadRequest(ctx, "group id is empty")
		return
	}

	start, end, err := helpers.ExtractIntervalFromQuery(ctx)
	if err != nil {
		helpers.HandleBadRequest(ctx, err.Error())
		return
	}

	lessons, err := e.manager.GetLessons(ctx, groupID, start, end)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, lessons)
}
