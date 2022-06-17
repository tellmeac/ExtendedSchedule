package faculty

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tellmeac/extended-schedule/pkg/faculty"
	"tellmeac/extended-schedule/pkg/handlers/helpers"
)

// New creates new endpoints for faculties.
func New(manager faculty.Manager) *Endpoints {
	return &Endpoints{manager: manager}
}

type Endpoints struct {
	manager faculty.Manager
}

func (e Endpoints) Bind(router gin.IRouter) {
	router.GET("faculties/", e.GetAllFaculties)
	router.GET("faculties/:facultyID/groups", e.GetFacultyGroups)
}

// GetAllFaculties - godoc
// @Router   /api/faculties [get]
// @Summary  Get all faculties
// @Tags     Faculty
// @Produce  application/json
// @Success  200  {array}  models.FacultyInfo
func (e Endpoints) GetAllFaculties(ctx *gin.Context) {
	faculties, err := e.manager.GetAllFaculties(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, faculties)
}

// GetFacultyGroups - godoc
// @Router   /api/faculties/{facultyID}/groups [get]
// @Summary  Get faculty's groups
// @Tags     Faculty
// @Param    facultyID  path  string  true  "Faculty ID"
// @Produce  application/json
// @Success  200  {array}  models.GroupInfo
// @Failure  404
func (e Endpoints) GetFacultyGroups(ctx *gin.Context) {
	facultyID := ctx.Param("facultyID")
	if facultyID == "" {
		helpers.HandleBadRequest(ctx, "faculty id is empty")
		return
	}

	groups, err := e.manager.GetByFaculty(ctx, facultyID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, groups)
}
