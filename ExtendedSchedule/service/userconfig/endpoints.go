package userconfig

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"tellmeac/extended-schedule/domain/aggregates"
	"tellmeac/extended-schedule/utils/shortcuts"
)

func NewEndpoints(service IService) *Endpoints {
	return &Endpoints{
		service: service,
	}
}

type Endpoints struct {
	service IService
}

func (e Endpoints) Bind(router gin.IRouter) {
	router.GET("/user/:userID/config", e.GetConfig)
	router.PATCH("/user/:userID/config", e.UpdateConfig)
}

func (e Endpoints) GetConfig(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("userID"))
	if err != nil {
		shortcuts.HandleBadRequest(ctx, err.Error())
		return
	}

	config, err := e.service.GetUserConfig(ctx, userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, config)
}

func (e Endpoints) UpdateConfig(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("userID"))
	if err != nil {
		shortcuts.HandleBadRequest(ctx, err.Error())
		return
	}

	var desired aggregates.UserConfig
	if err := ctx.ShouldBindJSON(&desired); err != nil {
		shortcuts.HandleBadRequest(ctx, err.Error())
		return
	}

	if userID != desired.UserID {
		shortcuts.HandleBadRequest(ctx, "userID in path should be equal to json body usedID field")
		return
	}

	if err := e.service.UpdateUserConfig(ctx, desired); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
