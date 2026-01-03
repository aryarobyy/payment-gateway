package model

import "time"

type User struct {
	ID        string    `json:"id" gorm:"type:uuid;primaryKey;"`
	Username  string    `json:"username" validate:"required" gorm:"unique"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password" validate:"required"`
	Role      Role      `json:"role" validate:"required"`
	LastLogin time.Time `json:"last_login" validate:"required"`

	CreatedAt time.Time `json:"created_at" gorm:"index"`
	UpdatedAt time.Time `json:"updated_at"`
}
