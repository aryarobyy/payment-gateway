package controller

import (
	"net/http"
	"time"

	"payment-gateway/internal/helper"
	"payment-gateway/internal/model"
	"payment-gateway/internal/service"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	service service.Service
}

func NewPaymentControllerl(service service.Service) *PaymentController {
	return &PaymentController{service: service}
}

func (h *PaymentController) Create(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.Payment()

	p := model.PostPayment{}

	if err := c.ShouldBindJSON(&p); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	payment, err := s.Create(ctx, p)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.Success(c, payment, "Payment created successfully")
}

func (h *PaymentController) GetMany(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.Payment()

	storeID := c.Query("store_id")
	if storeID == "" {
		helper.Error(c, http.StatusBadRequest, "Store ID parameter is required")
		return
	}

	limit, offset, err := helper.Pagination(c)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	payments, total, err := s.GetMany(ctx, storeID, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, gin.H{
		"data":  payments,
		"total": total,
	}, "Payments retrieved successfully")
}

func (h *PaymentController) GetByID(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.Payment()

	id := c.Param("id")
	if id == "" {
		helper.Error(c, http.StatusBadRequest, "ID parameter is required")
		return
	}

	payment, err := s.GetByID(ctx, id)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, payment, "Payment retrieved successfully")
}

func (h *PaymentController) GetByOrderID(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.Payment()

	orderID := c.Param("order_id")
	if orderID == "" {
		helper.Error(c, http.StatusBadRequest, "Order ID parameter is required")
		return
	}

	payment, err := s.GetByOrderID(ctx, orderID)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, payment, "Payment retrieved successfully")
}

func (h *PaymentController) GetByProviderRef(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.Payment()

	providerRef := c.Param("provider_ref")
	if providerRef == "" {
		helper.Error(c, http.StatusBadRequest, "Provider reference parameter is required")
		return
	}

	payment, err := s.GetByProviderRef(ctx, providerRef)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, payment, "Payment retrieved successfully")
}

func (h *PaymentController) GetByStatus(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.Payment()

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

	payments, total, err := s.GetByStatus(ctx, statusStr, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, gin.H{
		"data":  payments,
		"total": total,
	}, "Payments retrieved successfully")
}

func (h *PaymentController) UpdateStatus(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.Payment()

	paymentID := c.Param("id")
	if paymentID == "" {
		helper.Error(c, http.StatusBadRequest, "Payment ID parameter is required")
		return
	}

	var req struct {
		Status string `json:"status" validate:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := s.UpdateStatus(ctx, paymentID, req.Status); err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, nil, "Payment status updated successfully")
}

func (h *PaymentController) UpdateVerification(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.Payment()

	paymentID := c.Param("id")
	if paymentID == "" {
		helper.Error(c, http.StatusBadRequest, "Payment ID parameter is required")
		return
	}

	var req struct {
		VerifiedAt time.Time `json:"verified_at" validate:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := s.UpdateVerification(ctx, paymentID, req.VerifiedAt); err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, nil, "Payment verification updated successfully")
}
