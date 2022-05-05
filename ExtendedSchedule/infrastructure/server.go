package infrastructure

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	engine := gin.Default()
	engine.Use(cors.Default())
	return engine
}
