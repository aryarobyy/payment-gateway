package controller

import (
	"net/http"

	"payment-gateway/internal/helper"
	"payment-gateway/internal/helper/response"
	"payment-gateway/internal/model"
	"payment-gateway/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service service.Service
}

func NewAuthController(s service.Service) *AuthController {
	return &AuthController{service: s}
}

func (h *AuthController) Register(c *gin.Context) {
	ctx := c.Request.Context()
	authSrv := h.service.Auth()

	user := model.RegisterCredential{}

	if err := c.ShouldBindJSON(&user); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := authSrv.Register(ctx, user); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.Success(c, nil, "User registered successfully")
}

func (h *AuthController) Login(c *gin.Context) {
	ctx := c.Request.Context()
	authSrv := h.service.Auth()

	user := model.LoginCredential{}

	if err := c.ShouldBindJSON(&user); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	res, token, err := authSrv.Login(ctx, user)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.Success(c, response.ToUserResponse(*res), "User login successfully", token)
}

func (h *AuthController) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()
	authSrv := h.service.Auth()

	rToken, err := c.Cookie("token")
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	newToken, err := authSrv.RefreshToken(ctx, rToken)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.Success(c, nil, "Refresh token success", newToken)
}
