package controller

import (
	"net/http"

	"payment-gateway/internal/helper"
	"payment-gateway/internal/helper/response"
	"payment-gateway/internal/model"
	"payment-gateway/internal/service"

	"github.com/gin-gonic/gin"
)

type OrderItemController struct {
	service service.Service
}

func NewOrderItemController(service service.Service) *OrderItemController {
	return &OrderItemController{service: service}
}

func (h *OrderItemController) Create(c *gin.Context) {
	ctx := c.Request.Context()
	itemSrv := h.service.OrderItem()

	orderItem := model.OrderItem{}

	if err := c.ShouldBindJSON(&orderItem); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := itemSrv.Create(ctx, orderItem); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.Success(c, nil, "Order item created successfully")
}

func (h *OrderItemController) CreateBatch(c *gin.Context) {
	ctx := c.Request.Context()
	itemSrv := h.service.OrderItem()

	var orderItems []model.OrderItem

	if err := c.ShouldBindJSON(&orderItems); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := itemSrv.CreateBatch(ctx, orderItems); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.Success(c, nil, "Order items created successfully")
}

func (h *OrderItemController) GetMany(c *gin.Context) {
	ctx := c.Request.Context()
	itemSrv := h.service.OrderItem()

	orderID := c.Param("id")
	if orderID == "" {
		helper.Error(c, http.StatusBadRequest, "Order ID parameter is required")
		return
	}

	orderItems, err := itemSrv.GetMany(ctx, orderID)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, response.ToOrderItemResponseList(orderItems), "Order items retrieved successfully")
}

func (h *OrderItemController) GetByID(c *gin.Context) {
	ctx := c.Request.Context()
	itemSrv := h.service.OrderItem()

	id := c.Param("id")
	if id == "" {
		helper.Error(c, http.StatusBadRequest, "ID parameter is required")
		return
	}

	orderItem, err := itemSrv.GetByID(ctx, id)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, response.ToOrderItemResponse(*orderItem), "Order item retrieved successfully")
}
