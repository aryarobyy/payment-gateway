package routes

import (
	"payment-gateway/internal/controller"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, a controller.Controller) {
	routes := r.Group("/auth")
	auth := a.Auth()
	{
		routes.POST("/", auth.Register)
		routes.POST("/login", auth.Login)
		routes.POST("/refresh", auth.RefreshToken)
	}
}
