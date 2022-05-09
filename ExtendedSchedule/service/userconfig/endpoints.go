package userconfig

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tellmeac/extended-schedule/domain/aggregate"
	"tellmeac/extended-schedule/utils/middleware"
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
	router.GET("/user/config", e.GetConfig)
	router.PATCH("/user/config", e.UpdateConfig)
}

func (e Endpoints) GetConfig(ctx *gin.Context) {
	userIdentifier, err := middleware.GetGoogleEmail(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	config, err := e.service.GetUserConfig(ctx, userIdentifier)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, config)
}

func (e Endpoints) UpdateConfig(ctx *gin.Context) {
	var desired aggregate.UserConfig
	if err := ctx.ShouldBindJSON(&desired); err != nil {
		shortcuts.HandleBadRequest(ctx, err.Error())
		return
	}

	userIdentifier, err := middleware.GetGoogleEmail(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	if err := e.service.UpdateUserConfig(ctx, userIdentifier, desired); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
