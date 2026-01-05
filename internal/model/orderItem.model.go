package model

import "time"

type OrderItem struct {
	ID        string  `gorm:"type:char(36);primaryKey"`
	OrderID   string  `json:"order_id"`
	ProductID string  `json:"product_id"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	Subtotal  float64 `json:"subtotal"` // qty + price

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type UpdateOrderItem struct {
	OrderID   *string  `json:"order_id,omitempty"`
	ProductID *string  `json:"product_id,omitempty"`
	Price     *float64 `json:"price,omitempty"`
	Quantity  *int     `json:"quantity,omitempty"`
	Subtotal  *float64 `json:"subtotal,omitempty"` // qty + price

	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type OrderItemDTO struct {
	ProductID string  `json:"product_id"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	Subtotal  float64 `json:"subtotal"`
}
