package userconfig

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tellmeac/extended-schedule/common/models"
	"tellmeac/extended-schedule/lib/errors"
	"tellmeac/extended-schedule/lib/middleware"
	"tellmeac/extended-schedule/pkg/handlers/helpers"
	"tellmeac/extended-schedule/pkg/userconfig"
)

// New creates new endpoints for user schedule config.
func New(manager userconfig.Manager) *Endpoints {
	return &Endpoints{
		manager: manager,
	}
}

type Endpoints struct {
	manager userconfig.Manager
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
// @Success  200  {object}  models.UserConfig
func (e Endpoints) GetConfig(ctx *gin.Context) {
	email, err := middleware.GetGoogleEmail(ctx)
	if err != nil {
		errors.SendError(ctx, errors.ErrUnauthorized)
		return
	}

	userConfig, err := e.manager.GetByEmail(ctx, email)
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

	var desired models.UserConfig
	if err := ctx.ShouldBindJSON(&desired); err != nil {
		helpers.HandleBadRequest(ctx, err.Error())
		return
	}

	// prevent change others data.
	if email != desired.Email {
		errors.SendError(ctx, errors.ErrUnauthorized)
		return
	}

	if err := e.manager.Update(ctx, &desired); err != nil {
		errors.SendError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
