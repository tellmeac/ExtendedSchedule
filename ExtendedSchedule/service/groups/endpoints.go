package groups

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
}

func (e Endpoints) GetAllFaculties(ctx *gin.Context) {
	faculties, err := e.service.GetAllFaculties(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, faculties)
}

func (e Endpoints) GetFacultyGroups(ctx *gin.Context) {
	facultyID := ctx.Param("facultyID")

	groups, err := e.service.GetFacultyGroups(ctx, facultyID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, groups)
}
