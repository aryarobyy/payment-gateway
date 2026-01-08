package model

import "time"

type User struct {
	ID        string    `gorm:"type:char(36);primaryKey"`
	Username  string    `json:"username" validate:"required" gorm:"unique"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"-" validate:"required"`
	Role      Role      `json:"role" validate:"required"`
	LastLogin time.Time `json:"last_login" validate:"required"`

	CreatedAt time.Time `json:"created_at" gorm:"index"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RegisterCredential struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required"`
}

type LoginCredential struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type PasswordUpdate struct {
	UserID      string `json:"-"`
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}
