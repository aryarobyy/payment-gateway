package routes

import (
	"payment-gateway/internal/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, u controller.Controller) {
	routes := r.Group("/user")
	user := u.User()
	{
		routes.GET("/", user.GetMany)
		routes.GET("/:id", user.GetByID)
		routes.GET("/username/:username", user.GetByUsername)
		routes.GET("/role/:role", user.GetByRole)
	}
}
