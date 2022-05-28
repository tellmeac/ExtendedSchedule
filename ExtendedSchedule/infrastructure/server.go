package infrastructure

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewServer создает новый сервер.
func NewServer() *gin.Engine {
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowHeaders:     []string{"Authorization"},
	}))
	return engine
}
