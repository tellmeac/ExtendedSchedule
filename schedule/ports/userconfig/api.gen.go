// Package userconfig provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package userconfig

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// ExcludeRule defines model for ExcludeRule.
type ExcludeRule struct {
	lessonID string `json:"lessonId"`
	Pos      int    `json:"pos"`
}

// ExtendedGroup defines model for ExtendedGroup.
type ExtendedGroup struct {
	ID      string   `json:"id"`
	Lessons []Lesson `json:"lessons"`
}

// Lesson defines model for Lesson.
type Lesson struct {
	ID   string `json:"id"`
	Kind string `json:"kind"`
	Name string `json:"name"`
}

// Base group
type StudyGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Teacher based
type Teacher struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// UpdateConfig defines model for UpdateConfig.
type UpdateConfig struct {
	Base           *interface{}     `json:"base"`
	ExcludeRules   *[]ExcludeRule   `json:"excludeRules,omitempty"`
	ExtendedGroups *[]ExtendedGroup `json:"extendedGroups,omitempty"`
}

// UserConfig defines model for UserConfig.
type UserConfig struct {
	Base           *interface{}    `json:"base"`
	Email          string          `json:"email"`
	ExcludeRules   []ExcludeRule   `json:"excludeRules"`
	ExtendedGroups []ExtendedGroup `json:"extendedGroups"`
	ID             uuid.UUID       `json:"id"`
}

// GetUsersConfigParams defines parameters for GetUsersConfig.
type GetUsersConfigParams struct {
	Email openapi_types.Email `form:"email" json:"email"`
}

// PatchUsersConfigJSONBody defines parameters for PatchUsersConfig.
type PatchUsersConfigJSONBody = UpdateConfig

// PatchUsersConfigParams defines parameters for PatchUsersConfig.
type PatchUsersConfigParams struct {
	Email openapi_types.Email `form:"email" json:"email"`
}

// PatchUsersConfigJSONRequestBody defines body for PatchUsersConfig for application/json ContentType.
type PatchUsersConfigJSONRequestBody = PatchUsersConfigJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get user config
	// (GET /users/config)
	GetUsersConfig(c *gin.Context, params GetUsersConfigParams)
	// Update user config
	// (PATCH /users/config)
	PatchUsersConfig(c *gin.Context, params PatchUsersConfigParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(c *gin.Context)

// GetUsersConfig operation middleware
func (siw *ServerInterfaceWrapper) GetUsersConfig(c *gin.Context) {

	var err error

	c.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUsersConfigParams

	// ------------- Required query parameter "email" -------------
	if paramValue := c.Query("email"); paramValue != "" {

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Query argument email is required, but not found"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "email", c.Request.URL.Query(), &params.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter email: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetUsersConfig(c, params)
}

// PatchUsersConfig operation middleware
func (siw *ServerInterfaceWrapper) PatchUsersConfig(c *gin.Context) {

	var err error

	c.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PatchUsersConfigParams

	// ------------- Required query parameter "email" -------------
	if paramValue := c.Query("email"); paramValue != "" {

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Query argument email is required, but not found"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "email", c.Request.URL.Query(), &params.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter email: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PatchUsersConfig(c, params)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	router.GET(options.BaseURL+"/users/config", wrapper.GetUsersConfig)

	router.PATCH(options.BaseURL+"/users/config", wrapper.PatchUsersConfig)

	return router
}
