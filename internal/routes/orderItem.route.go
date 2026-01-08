package routes

import (
	"payment-gateway/internal/controller"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(r *gin.Engine, ctrl controller.Controller) {
	routes := r.Group("/order-item")
	orderItem := ctrl.OrderItem()
	{
		routes.POST("/", orderItem.Create)
		routes.POST("/batch", orderItem.CreateBatch)
		routes.GET("/:id", orderItem.GetByID)
		routes.GET("/order/:id", orderItem.GetMany)
	}
}

