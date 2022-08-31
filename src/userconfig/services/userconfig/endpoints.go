package userconfig

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tellmeac/extended-schedule/userconfig/domain/userconfig"
	"github.com/tellmeac/extended-schedule/userconfig/middle/errors"
	"github.com/tellmeac/extended-schedule/userconfig/pkg/middleware"
	"github.com/tellmeac/extended-schedule/userconfig/services/helpers"
)

// NewEndpoints creates new endpoints for user schedule config.
func NewEndpoints(manager Service) *Endpoints {
	return &Endpoints{
		service: manager,
	}
}

type Endpoints struct {
	service Service
}

func (e Endpoints) Bind(router gin.IRouter) {
	router.GET("/config", e.GetConfig)
	router.PATCH("/config", e.UpdateConfig)
}

// GetConfig - godoc
// @Router   /api/config [get]
// @Summary  Get config
// @Tags     Config
// @Produce  application/json
// @Param    Authorization  header  string  true  "Authorization bearer token"
// @Success  200  {object}  userconfig.UserConfig
func (e Endpoints) GetConfig(ctx *gin.Context) {
	email, err := middleware.GetGoogleEmail(ctx)
	if err != nil {
		errors.SendError(ctx, errors.ErrUnauthorized)
		return
	}

	userConfig, err := e.service.GetByEmail(ctx, email)
	if err != nil {
		errors.SendError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, userConfig)
}

// UpdateConfig - godoc
// @Router   /api/config [patch]
// @Summary  Get config
// @Tags     Config
// @Accept   application/json
// @Produce  application/json
// @Param    Authorization  header  string  true  "Authorization bearer token"
// @Param    desired  body  models.UserConfig  true  "Desired user config state"
// @Success  204
// @Failure  404
func (e Endpoints) UpdateConfig(ctx *gin.Context) {
	email, err := middleware.GetGoogleEmail(ctx)
	if err != nil {
		errors.SendError(ctx, errors.ErrUnauthorized)
		return
	}

	var desired userconfig.UserConfig
	if err := ctx.ShouldBindJSON(&desired); err != nil {
		helpers.HandleBadRequest(ctx, err.Error())
		return
	}

	// prevent changes to other people's user settings.
	if email != desired.Email {
		errors.SendError(ctx, errors.ErrUnauthorized)
		return
	}

	if err := e.service.Update(ctx, &desired); err != nil {
		errors.SendError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
