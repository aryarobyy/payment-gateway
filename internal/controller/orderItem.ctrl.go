package controller

import (
	"net/http"

	"payment-gateway/internal/helper"
	"payment-gateway/internal/model"
	"payment-gateway/internal/service"

	"github.com/gin-gonic/gin"
)

type OrderItemController struct {
	service service.Service
}

func NewOrderItemController(service service.Service) OrderItemController {
	return OrderItemController{service: service}
}

func (h *OrderItemController) Create(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.OrderItem()

	orderItem := model.OrderItem{}

	if err := c.ShouldBindJSON(&orderItem); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := s.Create(ctx, orderItem); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.Success(c, nil, "Order item created successfully")
}

func (h *OrderItemController) CreateBatch(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.OrderItem()

	var orderItems []model.OrderItem

	if err := c.ShouldBindJSON(&orderItems); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := s.CreateBatch(ctx, orderItems); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.Success(c, nil, "Order items created successfully")
}

func (h *OrderItemController) GetMany(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.OrderItem()

	orderID := c.Param("order_id")
	if orderID == "" {
		helper.Error(c, http.StatusBadRequest, "Order ID parameter is required")
		return
	}

	orderItems, err := s.GetMany(ctx, orderID)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, orderItems, "Order items retrieved successfully")
}

func (h *OrderItemController) GetByID(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.OrderItem()

	id := c.Param("id")
	if id == "" {
		helper.Error(c, http.StatusBadRequest, "ID parameter is required")
		return
	}

	orderItem, err := s.GetByID(ctx, id)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, orderItem, "Order item retrieved successfully")
}