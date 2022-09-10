// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package ports

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get faculty list
	// (GET /faculties)
	GetFaculties(c *gin.Context)
	// Get group's lesson list
	// (GET /lessons/byGroup/{id})
	GetLessonsByGroupId(c *gin.Context, id string)
	// Get group's schedule
	// (GET /schedule/byGroup/{id})
	GetScheduleByGroupId(c *gin.Context, id string)
	// Get personal schedule
	// (GET /users/{id}/schedule)
	GetUsersIdSchedule(c *gin.Context, id uuid.UUID)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(c *gin.Context)

// GetFaculties operation middleware
func (siw *ServerInterfaceWrapper) GetFaculties(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetFaculties(c)
}

// GetLessonsByGroupId operation middleware
func (siw *ServerInterfaceWrapper) GetLessonsByGroupId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetLessonsByGroupId(c, id)
}

// GetScheduleByGroupId operation middleware
func (siw *ServerInterfaceWrapper) GetScheduleByGroupId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetScheduleByGroupId(c, id)
}

// GetUsersIdSchedule operation middleware
func (siw *ServerInterfaceWrapper) GetUsersIdSchedule(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id uuid.UUID

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter id: %s", err)})
		return
	}

	c.Set(BearerAuthScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetUsersIdSchedule(c, id)
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

	router.GET(options.BaseURL+"/faculties", wrapper.GetFaculties)

	router.GET(options.BaseURL+"/lessons/byGroup/:id", wrapper.GetLessonsByGroupId)

	router.GET(options.BaseURL+"/schedule/byGroup/:id", wrapper.GetScheduleByGroupId)

	router.GET(options.BaseURL+"/users/:id/schedule", wrapper.GetUsersIdSchedule)

	return router
}