package infrastructure

import (
	"tellmeac/extended-schedule/config"
	"tellmeac/extended-schedule/utils/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewServer создает новый сервер.
func NewServer(cfg config.Config) *gin.Engine {
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowHeaders:     []string{"Authorization"},
	}))

	engine.Use(middleware.GoogleOAuth2(cfg.IsDebug))

	return engine
}
