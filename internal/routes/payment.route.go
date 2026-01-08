package routes

import (
	"payment-gateway/internal/controller"
	"payment-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.Engine, ctrl controller.Controller) {
	routes := r.Group("/payment")
	payment := ctrl.Payment()
	{
		routes.POST("/", payment.Create)
		routes.GET("/", payment.GetMany)
		routes.GET("/:id", payment.GetByID)
		routes.GET("/order/:id", payment.GetByOrderID)
		routes.GET("/provider/:ref", payment.GetByProviderRef)
		routes.GET("/status/:status", payment.GetByStatus)
		routes.PUT("/:id/status", middleware.AuthMiddleware(), payment.UpdateStatus)
		routes.PUT("/:id/verification", payment.UpdateVerification)
	}
}
