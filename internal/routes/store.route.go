package routes

import (
	"payment-gateway/internal/controller"
	"payment-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func StoreRoutes(r *gin.Engine, ctrl controller.Controller) {
	routes := r.Group("/store")
	store := ctrl.Store()
	{
		routes.POST("/", store.Create)
		routes.GET("/", store.GetMany)
		routes.GET("/:id", store.GetByID)
		routes.GET("/owner/:id", store.GetByOwnerID)
		routes.GET("/active/:is_active", store.GetByIsActive)
		routes.PUT("/:id", middleware.AuthMiddleware(), store.Update)
	}
}
