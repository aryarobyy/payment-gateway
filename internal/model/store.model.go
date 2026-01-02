package model

import "time"

type Store struct {
	ID          string  `json:"id" gorm:"type:uuid;primaryKey;"`
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	IsActive    bool    `json:"is_active" gorm:"default:false"`
	OwnerID     string  `json:"owner_id"`

	CreatedAt time.Time `json:"created_at" gorm:"index"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateStore struct {
	Name        *string `json:"name,omitempty" validate:"required"`
	Description *string `json:"description,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty" gorm:"default:false"`
	OwnerID     *string `json:"owner_id,omitempty"`

	UpdatedAt time.Time `json:"updated_at"`
}
