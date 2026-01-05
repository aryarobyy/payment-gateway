package controller

import (
	"net/http"

	"payment-gateway/internal/helper"
	"payment-gateway/internal/model"
	"payment-gateway/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service service.Service
}

func NewProductController(service service.Service) ProductController {
	return ProductController{service: service}
}

func (h *ProductController) Create(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.Product()

	p := model.PostProduct{}

	if err := c.ShouldBindJSON(&p); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := s.Create(ctx, p); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.Success(c, nil, "Product created successfully")
}

func (h *ProductController) GetMany(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.Product()

	storeID := c.Query("store_id")

	limit, offset, err := helper.Pagination(c)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	products, total, err := s.GetMany(ctx, storeID, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, gin.H{
		"data":  products,
		"total": total,
	}, "Products retrieved successfully")
}

func (h *ProductController) GetByCategory(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.Product()

	category := c.Param("category")
	if category == "" {
		helper.Error(c, http.StatusBadRequest, "Category parameter is required")
		return
	}

	limit, offset, err := helper.Pagination(c)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	products, total, err := s.GetByCategory(ctx, category, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, gin.H{
		"data":  products,
		"total": total,
	}, "Products retrieved successfully")
}

func (h *ProductController) GetByID(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.Product()

	id := c.Param("id")
	if id == "" {
		helper.Error(c, http.StatusBadRequest, "ID parameter is required")
		return
	}

	product, err := s.GetByID(ctx, id)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, product, "Product retrieved successfully")
}

func (h *ProductController) GetByActive(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.Product()

	isActiveStr := c.Param("is_active")
	if isActiveStr == "" {
		helper.Error(c, http.StatusBadRequest, "isActive parameter is required")
		return
	}

	limit, offset, err := helper.Pagination(c)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	products, total, err := s.GetByActive(ctx, isActiveStr, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, gin.H{
		"data":  products,
		"total": total,
	}, "Products retrieved successfully")
}

func (h *ProductController) Update(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.Product()

	id := c.Param("id")
	if id == "" {
		helper.Error(c, http.StatusBadRequest, "ID parameter is required")
		return
	}

	updateProduct := model.UpdateProduct{}

	if err := c.ShouldBindJSON(&updateProduct); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	product, err := s.Update(ctx, updateProduct, id)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, product, "Product updated successfully")
}
