package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Endpoints E

type E struct{}

func (e E) Bind(router gin.IRouter) {
	router.GET("/health", e.Health)
}

func (e E) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Status": "OK"})
}
