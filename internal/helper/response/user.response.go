package response

import (
	"payment-gateway/internal/model"
	"time"
)

type UserResponse struct {
	ID        string     `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Role      model.Role `json:"role"`
	LastLogin time.Time  `json:"last_login"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func ToUserResponse(u model.User) UserResponse {
	return UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Role:      u.Role,
		LastLogin: u.LastLogin,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToUserResponseList(users []model.User) []UserResponse {
	responses := make([]UserResponse, len(users))
	for i, u := range users {
		responses[i] = ToUserResponse(u)
	}
	return responses
}
