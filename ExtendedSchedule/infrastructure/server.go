package infrastructure

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"tellmeac/extended-schedule/config"
	"tellmeac/extended-schedule/middle/middleware"
)

// NewServer creates new http server.
func NewServer(cfg config.Config) *gin.Engine {
	engine := gin.Default()

	// enable CORS as debug only option
	if cfg.IsDebug {
		log.Debug().Msg("Use cors policy middleware")
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

	// handle unauthorized access to API
	engine.Use(middleware.GoogleOAuth2(cfg.IsDebug))

	return engine
}
