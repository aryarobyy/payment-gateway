package model

import "time"

type Product struct {
	ID       string  `gorm:"type:char(36);primaryKey"`
	StoreID  string  `json:"store_id" validate:"required"`
	Name     string  `json:"name" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
	IsActive bool    `json:"is_active" gorm:"default:false"`
	Category string  `json:"category"`

	CreatedAt time.Time `json:"created_at" gorm:"index"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostProduct struct {
	ID       string  `gorm:"type:char(36);primaryKey"`
	StoreID  string  `json:"store_id" validate:"required"`
	Name     string  `json:"name" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
	IsActive bool    `json:"is_active" gorm:"default:false"`
	Category string  `json:"category"`

	CreatedAt time.Time `json:"created_at"`
}

type UpdateProduct struct {
	StoreID  *string  `json:"store_id,omitempty" validate:"required"`
	Name     *string  `json:"name,omitempty" validate:"required"`
	Price    *float64 `json:"price,omitempty" validate:"required"`
	IsActive *bool    `json:"is_active,omitempty" gorm:"default:false"`
	Category *string  `json:"category,omitempty"`

	UpdatedAt time.Time `json:"updated_at"`
}
