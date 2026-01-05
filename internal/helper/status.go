package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Token   string      `json:"token,omitempty"`
}

func Success(c *gin.Context, data interface{}, message string, token ...string) {
	resp := Response{
		Code:    http.StatusOK,
		Status:  "success",
		Message: message,
		Data:    data,
	}

	if len(token) > 0 {
		resp.Token = token[0]
	}

	c.JSON(http.StatusOK, resp)
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Status:  "error",
		Message: message,
	})
}
