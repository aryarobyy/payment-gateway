package middleware

//
// import (
// 	"fmt"
// 	"net/http"
// 	"strings"
//
// 	"github.com/gin-gonic/gin"
// )
//
// type AuthMiddleware struct {
// 	// Add any dependencies needed for authentication
// 	// e.g., JWT secret, user repository, etc.
// }
//
// func NewAuthMiddleware() *AuthMiddleware {
// 	return &AuthMiddleware{}
// }
//
// func (m *AuthMiddleware) JWTAuth() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
// 			c.Abort()
// 			return
// 		}
//
// 		// Check if the header has the correct format "Bearer {token}"
// 		parts := strings.Split(authHeader, " ")
// 		if len(parts) != 2 || parts[0] != "Bearer" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
// 			c.Abort()
// 			return
// 		}
//
// 		tokenString := parts[1]
//
// 		// Here you would normally validate the JWT token
// 		// For now, I'll simulate a successful validation and set a user ID
// 		// In a real implementation, you would extract the user ID from the token
// 		userID := uint(1) // This would come from the token in a real implementation
//
// 		// Set the user ID in the context for use in handlers
// 		c.Set("user_id", userID)
//
// 		// Continue to the next handler
// 		c.Next()
// 	}
// }
