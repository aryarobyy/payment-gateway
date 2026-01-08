package helper

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"payment-gateway/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ParseExpiry(s string) (time.Duration, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, errors.New("empty expiry string")
	}

	if strings.HasSuffix(s, "d") {
		nStr := strings.TrimSuffix(s, "d")
		n, err := strconv.Atoi(nStr)
		if err != nil {
			return 0, err
		}
		return time.Duration(n) * 24 * time.Hour, nil
	}

	d, err := time.ParseDuration(s)
	if err != nil {
		return 0, err
	}
	return d, nil
}

func GenerateAccessToken(user *model.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	expiryStr := os.Getenv("JWT_EXPIRED")
	if secret == "" {
		return "", errors.New("JWT_SECRET is not set in environment")
	}

	if expiryStr == "" {
		expiryStr = "30m"
	}

	duration, err := ParseExpiry(expiryStr)
	if err != nil {
		return "", err
	}

	var userEmail string

	if user.Email != "" {
		userEmail = user.Email
	} else {
		userEmail = ""
	}

	expireAt := time.Now().Add(duration)
	claims := model.ClaimsModel{
		UserID: user.ID,
		Role:   user.Role,
		Email:  userEmail,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expireAt),
			Subject:   user.ID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func GenerateRefreshToken(c *gin.Context, user *model.User) (string, error) {
	secret := os.Getenv("JWT_REFRESH_SECRET")
	if secret == "" {
		secret = os.Getenv("JWT_SECRET")
	}

	expiryStr := os.Getenv("JWT_REFRESH_EXPIRED")
	if expiryStr == "" {
		expiryStr = "7d"
	}

	duration, err := ParseExpiry(expiryStr)
	if err != nil {
		return "", err
	}

	expireAt := time.Now().Add(duration)
	claims := jwt.RegisteredClaims{
		Subject:   user.ID,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(expireAt),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	c.SetCookie(
		"token",
		ss,
		int(duration.Seconds()),
		"/",
		"",
		false,
		true,
	)

	return ss, nil
}

func ParseAndValidateToken(tokenString string) (*model.ClaimsModel, error) {
	secret := os.Getenv("JWT_SECRET")

	token, err := jwt.ParseWithClaims(tokenString, &model.ClaimsModel{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*model.ClaimsModel)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func ValidateRefreshToken(c *gin.Context) (string, error) {
	secret := os.Getenv("JWT_REFRESH_SECRET")
	if secret == "" {
		secret = os.Getenv("JWT_SECRET")
	}

	tokenString, err := c.Cookie("token")
	if err != nil {
		return "", err
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Method.Alg())
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", fmt.Errorf("error parsing refresh token: %w", err)
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid or expired refresh token")
	}

	if claims.Subject == "" {
		return "", errors.New("missing subject in refresh token")
	}

	return claims.Subject, nil
}
