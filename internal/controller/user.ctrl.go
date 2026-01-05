package controller

import (
	"net/http"

	"payment-gateway/internal/helper"
	"payment-gateway/internal/model"
	"payment-gateway/internal/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.Service
}

func NewUserController(s service.Service) *UserController {
	return &UserController{service: s}
}

func (h *UserController) Register(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.User()

	user := model.RegisterCredential{}

	if err := c.ShouldBindJSON(&user); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := s.Register(ctx, user); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.Success(c, nil, "User registered successfully")
}

func (h *UserController) Login(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.User()

	user := model.LoginCredential{}

	if err := c.ShouldBindJSON(&user); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	res, _, err := s.Login(ctx, user) // token disable dlu
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.Success(c, res, "User login successfully")
}

func (h *UserController) GetMany(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.User()

	limit, offset, err := helper.Pagination(c)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	users, total, err := s.GetMany(ctx, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, gin.H{
		"data":  users,
		"total": total,
	}, "Users retrieved successfully")
}

func (h *UserController) GetByRole(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.User()

	role := c.Param("role")
	if role == "" {
		helper.Error(c, http.StatusBadRequest, "Role parameter is required")
		return
	}

	limit, offset, err := helper.Pagination(c)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	users, total, err := s.GetByRole(ctx, role, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, gin.H{
		"data":  users,
		"total": total,
	}, "Users retrieved successfully")
}

func (h *UserController) GetByID(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.User()

	id := c.Param("id")
	if id == "" {
		helper.Error(c, http.StatusBadRequest, "ID parameter is required")
		return
	}

	user, err := s.GetByID(ctx, id)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, user, "User retrieved successfully")
}

func (h *UserController) GetByUsername(c *gin.Context) {
	ctx := c.Request.Context()
	s := h.service.User()

	username := c.Param("username")
	if username == "" {
		helper.Error(c, http.StatusBadRequest, "Username parameter is required")
		return
	}

	user, err := s.GetByUsername(ctx, username)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, user, "User retrieved successfully")
}
