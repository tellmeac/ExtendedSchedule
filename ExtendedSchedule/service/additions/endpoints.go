package additions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewEndpoints(service IService) *Endpoints {
	return &Endpoints{service: service}
}

type Endpoints struct {
	service IService
}

func (e Endpoints) Bind(router gin.IRouter) {
	router.GET("faculties/", e.GetAllFaculties)
	router.GET("faculties/:facultyID/groups", e.GetFacultyGroups)
	router.GET("lessons/", e.GetLessonList)
}

// GetAllFaculties - godoc
// @Router   /api/faculties [get]
// @Summary  Get all faculties
// @Tags     Groups
// @Produce  application/json
// @Success  200  {array}  entity.FacultyInfo
// @Failure  500
func (e Endpoints) GetAllFaculties(ctx *gin.Context) {
	faculties, err := e.service.GetAllFaculties(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, faculties)
}

// GetFacultyGroups - godoc
// @Router   /api/faculties/{facultyID}/groups [get]
// @Summary  Get faculty's groups
// @Tags     Groups
// @Param facultyID path string true "Faculty ID"
// @Produce  application/json
// @Success  200  {array}  entity.GroupInfo
// @Failure  500
func (e Endpoints) GetFacultyGroups(ctx *gin.Context) {
	facultyID := ctx.Param("facultyID")

	groups, err := e.service.GetFacultyGroups(ctx, facultyID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, groups)
}

// GetLessonList - godoc
// @Router   /api/lessons [get]
// @Summary  Get group's lesson list
// @Tags     Groups
// @Produce  application/json
// @Success  200  {array}  provider.LessonInfo
// @Failure  500
func (e Endpoints) GetLessonList(ctx *gin.Context) {
	groupID := ctx.Query("groupId")

	lessons, err := e.service.GetGroupLessons(ctx, groupID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, lessons)
}
