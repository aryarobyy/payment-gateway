package controller

import (
	"errors"
	"net/http"
	"strings"

	"payment-gateway/internal/helper"
	"payment-gateway/internal/helper/response"
	"payment-gateway/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	service service.Service
}

func NewUserController(s service.Service) *UserController {
	return &UserController{service: s}
}

func (h *UserController) GetMany(c *gin.Context) {
	ctx := c.Request.Context()
	authSrv := h.service.User()

	limit, offset, err := helper.Pagination(c)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	users, total, err := authSrv.GetMany(ctx, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, gin.H{
		"data":  response.ToUserResponseList(users),
		"total": total,
	}, "Users retrieved successfully")
}

func (h *UserController) GetByRole(c *gin.Context) {
	ctx := c.Request.Context()
	authSrv := h.service.User()

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
	users, total, err := authSrv.GetByRole(ctx, role, limit, offset)
	if err != nil {
		if strings.Contains(err.Error(), "invalid role") {
			helper.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, gin.H{
		"data":  response.ToUserResponseList(users),
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Error(c, http.StatusNotFound, "User not found")
			return
		}
		if strings.Contains(err.Error(), "invalid UUID") {
			helper.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, response.ToUserResponse(*user), "User retrieved successfully")
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Error(c, http.StatusNotFound, "User not found")
			return
		}
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, response.ToUserResponse(*user), "User retrieved successfully")
}
