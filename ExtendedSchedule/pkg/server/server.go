package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"tellmeac/extended-schedule/lib/middleware"
	"tellmeac/extended-schedule/pkg/config"
)

// New creates new http server.
func New(cfg config.Config) *gin.Engine {
	engine := gin.Default()

	// enable cors for local debugging
	if cfg.IsDebug {
		engine.Use(cors.New(cors.Config{
			AllowAllOrigins:  true,
			AllowCredentials: true,
			AllowHeaders: []string{
				"Authorization",
				"Content-Type",
			},
			AllowMethods: []string{"PATCH", "GET"},
		}))
	}

	// handle unauthorized access
	engine.Use(middleware.GoogleOAuth2(cfg.IsDebug))

	return engine
}
