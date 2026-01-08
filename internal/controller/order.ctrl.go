package controller

import (
	"net/http"

	"payment-gateway/internal/helper"
	"payment-gateway/internal/helper/response"
	"payment-gateway/internal/model"
	"payment-gateway/internal/service"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	service service.Service
}

func NewOrderController(service service.Service) *OrderController {
	return &OrderController{service: service}
}

func (h *OrderController) Create(c *gin.Context) {
	ctx := c.Request.Context()
	orderSrv := h.service.Order()

	order := model.Order{}

	if err := c.ShouldBindJSON(&order); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := orderSrv.Create(ctx, order); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.Success(c, nil, "Order created successfully")
}

func (h *OrderController) GetMany(c *gin.Context) {
	ctx := c.Request.Context()
	orderSrv := h.service.Order()

	limit, offset, err := helper.Pagination(c)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	orders, total, err := orderSrv.GetMany(ctx, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, gin.H{
		"data":  response.ToOrderResponseList(orders),
		"total": total,
	}, "Orders retrieved successfully")
}

func (h *OrderController) GetByStoreID(c *gin.Context) {
	ctx := c.Request.Context()
	orderSrv := h.service.Order()

	storeID := c.Param("id")
	if storeID == "" {
		helper.Error(c, http.StatusBadRequest, "Store ID parameter is required")
		return
	}

	limit, offset, err := helper.Pagination(c)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	orders, total, err := orderSrv.GetByStoreID(ctx, storeID, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, gin.H{
		"data":  response.ToOrderResponseList(orders),
		"total": total,
	}, "Orders retrieved successfully")
}

func (h *OrderController) GetByStatus(c *gin.Context) {
	ctx := c.Request.Context()
	orderSrv := h.service.Order()

	statusStr := c.Param("status")
	if statusStr == "" {
		helper.Error(c, http.StatusBadRequest, "Status parameter is required")
		return
	}

	limit, offset, err := helper.Pagination(c)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	orders, total, err := orderSrv.GetByStatus(ctx, statusStr, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, gin.H{
		"data":  response.ToOrderResponseList(orders),
		"total": total,
	}, "Orders retrieved successfully")
}

func (h *OrderController) GetByID(c *gin.Context) {
	ctx := c.Request.Context()
	orderSrv := h.service.Order()

	id := c.Param("id")
	if id == "" {
		helper.Error(c, http.StatusBadRequest, "ID parameter is required")
		return
	}

	order, err := orderSrv.GetByID(ctx, id)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, response.ToOrderResponse(*order), "Order retrieved successfully")
}

func (h *OrderController) UpdateStatus(c *gin.Context) {
	ctx := c.Request.Context()
	orderSrv := h.service.Order()

	orderID := c.Param("id")
	if orderID == "" {
		helper.Error(c, http.StatusBadRequest, "Order ID parameter is required")
		return
	}

	var req struct {
		Status model.Status `json:"status" validate:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := orderSrv.UpdateStatus(ctx, orderID, string(req.Status)); err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, nil, "Order status updated successfully")
}
