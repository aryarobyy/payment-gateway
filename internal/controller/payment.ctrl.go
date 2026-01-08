package controller

import (
	"net/http"
	"time"

	"payment-gateway/internal/helper"
	"payment-gateway/internal/helper/response"
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
	paymentSrv := h.service.Payment()

	p := model.PostPayment{}

	if err := c.ShouldBindJSON(&p); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	payment, err := paymentSrv.Create(ctx, p)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.Success(c, response.ToPaymentResponse(*payment), "Payment created successfully")
}

func (h *PaymentController) GetMany(c *gin.Context) {
	ctx := c.Request.Context()
	paymentSrv := h.service.Payment()

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

	payments, total, err := paymentSrv.GetMany(ctx, storeID, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, gin.H{
		"data":  response.ToPaymentResponseList(payments),
		"total": total,
	}, "Payments retrieved successfully")
}

func (h *PaymentController) GetByID(c *gin.Context) {
	ctx := c.Request.Context()
	paymentSrv := h.service.Payment()

	id := c.Param("id")
	if id == "" {
		helper.Error(c, http.StatusBadRequest, "ID parameter is required")
		return
	}

	payment, err := paymentSrv.GetByID(ctx, id)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, response.ToPaymentResponse(*payment), "Payment retrieved successfully")
}

func (h *PaymentController) GetByOrderID(c *gin.Context) {
	ctx := c.Request.Context()
	paymentSrv := h.service.Payment()

	orderID := c.Param("id")
	if orderID == "" {
		helper.Error(c, http.StatusBadRequest, "Order ID parameter is required")
		return
	}

	payment, err := paymentSrv.GetByOrderID(ctx, orderID)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, response.ToPaymentResponse(*payment), "Payment retrieved successfully")
}

func (h *PaymentController) GetByProviderRef(c *gin.Context) {
	ctx := c.Request.Context()
	paymentSrv := h.service.Payment()

	providerRef := c.Param("ref")
	if providerRef == "" {
		helper.Error(c, http.StatusBadRequest, "Provider reference parameter is required")
		return
	}

	payment, err := paymentSrv.GetByProviderRef(ctx, providerRef)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, response.ToPaymentResponse(*payment), "Payment retrieved successfully")
}

func (h *PaymentController) GetByStatus(c *gin.Context) {
	ctx := c.Request.Context()
	paymentSrv := h.service.Payment()

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

	payments, total, err := paymentSrv.GetByStatus(ctx, statusStr, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, gin.H{
		"data":  response.ToPaymentResponseList(payments),
		"total": total,
	}, "Payments retrieved successfully")
}

func (h *PaymentController) UpdateStatus(c *gin.Context) {
	ctx := c.Request.Context()
	paymentSrv := h.service.Payment()

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

	if err := paymentSrv.UpdateStatus(ctx, paymentID, req.Status); err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, nil, "Payment status updated successfully")
}

func (h *PaymentController) UpdateVerification(c *gin.Context) {
	ctx := c.Request.Context()
	paymentSrv := h.service.Payment()

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

	if err := paymentSrv.UpdateVerification(ctx, paymentID, req.VerifiedAt); err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, nil, "Payment verification updated successfully")
}
