package middleware

import (
	"net/http"
	"strings"

	"payment-gateway/internal/helper"
	"payment-gateway/internal/model"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		claims, err := helper.ParseAndValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		c.Set("user", claims)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func RoleCheck(role ...model.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		val, exists := c.Get("role")

		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: role info missing"})
			return
		}

		userRole, ok := val.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "system error: invalid role type"})
			return
		}

		allowed := false
		for _, r := range role {
			if strings.EqualFold(string(r), userRole) {
				allowed = true
				break
			}
		}

		if !allowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden: insufficient permissions"})
			return
		}

		c.Next()
	}
}
