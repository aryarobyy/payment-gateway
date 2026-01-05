package routes

import (
	"payment-gateway/internal/controller"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine, ctrl controller.Controller) {
	routes := r.Group("/order")
	order := ctrl.Order()
	{
		routes.POST("/", order.Create)
		routes.GET("/", order.GetMany)
		routes.GET("/:id", order.GetByID)
		routes.GET("/store/:store_id", order.GetByStoreID)
		routes.GET("/status/:status", order.GetByStatus)
		routes.PUT("/:id/status", order.UpdateStatus)
	}
}