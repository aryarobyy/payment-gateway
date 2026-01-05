package routes

import (
	"payment-gateway/internal/controller"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, ctrl controller.Controller) {
	routes := r.Group("/product")
	product := ctrl.Product()
	{
		routes.POST("/", product.Create)
		routes.GET("/", product.GetMany)
		routes.GET("/:id", product.GetByID)
		routes.GET("/category/:category", product.GetByCategory)
		routes.GET("/active/:is_active", product.GetByActive)
		routes.PUT("/:id", product.Update)
	}
}