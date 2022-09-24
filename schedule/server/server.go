package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"tellmeac/extended-schedule/config"
)

// NewServer creates new http server.
func NewServer() *gin.Engine {
	engine := gin.Default()

	cfg := config.Get()

	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// enable CORS as debug only option
	if cfg.Debug {
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
	//engine.Use(middleware.GoogleOAuth2(cfg.Debug))

	return engine
}
