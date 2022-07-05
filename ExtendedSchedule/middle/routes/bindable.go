package routes

import "github.com/gin-gonic/gin"

// Bindable represents bindable endpoint.
type Bindable interface {
	Bind(router gin.IRouter)
}
