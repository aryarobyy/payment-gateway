package routes

import (
	"payment-gateway/internal/controller"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.Engine, ctrl controller.Controller) {
	routes := r.Group("/payment")
	payment := ctrl.Payment()
	{
		routes.POST("/", payment.Create)
		routes.GET("/", payment.GetMany)
		routes.GET("/:id", payment.GetByID)
		routes.GET("/order/:order_id", payment.GetByOrderID)
		routes.GET("/provider/:provider_ref", payment.GetByProviderRef)
		routes.GET("/status/:status", payment.GetByStatus)
		routes.PUT("/:id/status", payment.UpdateStatus)
		routes.PUT("/:id/verification", payment.UpdateVerification)
	}
}

