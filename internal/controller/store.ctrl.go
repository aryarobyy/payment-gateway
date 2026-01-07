package controller

import (
	"net/http"

	"payment-gateway/internal/helper"
	"payment-gateway/internal/helper/response"
	"payment-gateway/internal/model"
	"payment-gateway/internal/service"

	"github.com/gin-gonic/gin"
)

type StoreController struct {
	service service.Service
}

func NewStoreController(service service.Service) *StoreController {
	return &StoreController{service: service}
}

func (h *StoreController) Create(c *gin.Context) {
	ctx := c.Request.Context()
	storeSrv := h.service.Store()

	s := model.Store{}

	if err := c.ShouldBindJSON(&s); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := storeSrv.Create(ctx, s); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.Success(c, nil, "Store created successfully")
}

func (h *StoreController) GetMany(c *gin.Context) {
	ctx := c.Request.Context()
	storeSrv := h.service.Store()

	limit, offset, err := helper.Pagination(c)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	stores, total, err := storeSrv.GetMany(ctx, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, gin.H{
		"data":  response.ToStoreResponseList(stores),
		"total": total,
	}, "Stores retrieved successfully")
}

func (h *StoreController) GetByOwnerID(c *gin.Context) {
	ctx := c.Request.Context()
	storeSrv := h.service.Store()

	ownerID := c.Param("id")
	if ownerID == "" {
		helper.Error(c, http.StatusBadRequest, "Owner ID parameter is required")
		return
	}

	limit, offset, err := helper.Pagination(c)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	stores, total, err := storeSrv.GetByOwnerID(ctx, ownerID, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, gin.H{
		"data":  response.ToStoreResponseList(stores),
		"total": total,
	}, "Stores retrieved successfully")
}

func (h *StoreController) GetByID(c *gin.Context) {
	ctx := c.Request.Context()
	storeSrv := h.service.Store()

	id := c.Param("id")
	if id == "" {
		helper.Error(c, http.StatusBadRequest, "ID parameter is required")
		return
	}

	store, err := storeSrv.GetByID(ctx, id)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, response.ToStoreResponse(*store), "Store retrieved successfully")
}

func (h *StoreController) GetByIsActive(c *gin.Context) {
	ctx := c.Request.Context()
	storeSrv := h.service.Store()

	isActiveStr := c.Param("is_active")
	if isActiveStr == "" {
		helper.Error(c, http.StatusBadRequest, "isActive parameter is required")
		return
	}

	// Simple conversion since service expects bool.
	// In ProductController it passed the string to service?
	// Let's check service again.
	// ProductService.GetByActive takes string? No, StoreService.GetByIsActive takes bool.
	// Let's check ProductService.GetByActive signature in `product.srv.go`?
	// I don't have product.srv.go content, but I saw product.ctrl.go call it with string.
	// Wait, let's look at StoreService.GetByIsActive signature.
	// func (s *storeService) GetByIsActive(ctx context.Context, isActive bool, limit int, offset int)

	isActive := isActiveStr == "true" || isActiveStr == "1"

	limit, offset, err := helper.Pagination(c)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	stores, total, err := storeSrv.GetByIsActive(ctx, isActive, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, gin.H{
		"data":  response.ToStoreResponseList(stores),
		"total": total,
	}, "Stores retrieved successfully")
}

func (h *StoreController) Update(c *gin.Context) {
	ctx := c.Request.Context()
	storeSrv := h.service.Store()

	id := c.Param("id")
	if id == "" {
		helper.Error(c, http.StatusBadRequest, "ID parameter is required")
		return
	}

	updateStore := model.UpdateStore{}

	if err := c.ShouldBindJSON(&updateStore); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	store, err := storeSrv.Update(ctx, updateStore, id)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, response.ToStoreResponse(*store), "Store updated successfully")
}
