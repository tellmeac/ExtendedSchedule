package shortcuts

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleBadRequest(ctx *gin.Context, errorMsg string) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errorMsg})
}
