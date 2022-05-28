package userconfig

import (
	"net/http"
	"tellmeac/extended-schedule/domain/aggregate"
	"tellmeac/extended-schedule/utils/middleware"
	"tellmeac/extended-schedule/utils/shortcuts"

	"github.com/gin-gonic/gin"
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

// GetConfig - godoc
// @Router   /api/user/config [get]
// @Summary  Get config
// @Tags     Config
// @Produce  application/json
// @Success  200  {object}  aggregate.UserConfig
// @Failure  500
func (e Endpoints) GetConfig(ctx *gin.Context) {
	email, err := middleware.GetGoogleEmail(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	config, err := e.service.GetUserConfig(ctx, email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, config)
}

// UpdateConfig - godoc
// @Router   /api/user/config [patch]
// @Summary  Get config
// @Tags     Config
// @Accept   application/json
// @Produce  application/json
// @Param    desired  body  aggregate.UserConfig  true  "Desired user config state"
// @Success  204
// @Failure  500
func (e Endpoints) UpdateConfig(ctx *gin.Context) {
	email, err := middleware.GetGoogleEmail(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	var desired aggregate.UserConfig
	if err := ctx.ShouldBindJSON(&desired); err != nil {
		shortcuts.HandleBadRequest(ctx, err.Error())
		return
	}

	if email != desired.Email {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := e.service.UpdateUserConfig(ctx, desired); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
