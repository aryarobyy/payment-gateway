package model

import "time"

type Product struct {
	ID       string  `json:"id" gorm:"type:uuid;primaryKey;"`
	StoreID  string  `json:"store_id" validate:"required"`
	Name     string  `json:"name" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
	IsActive bool    `json:"is_active" gorm:"default:false"`

	CreatedAt time.Time `json:"created_at" gorm:"index"`
	UpdatedAt time.Time `json:"updated_at"`
}
