package model

import "time"

type Order struct {
	ID          string     `json:"id" gorm:"type:uuid;primaryKey"`
	StoreID     string     `json:"store_id" validate:"required"`
	Status      Status     `json:"status" validate:"required"`
	TableID     string     `json:"table_id" validate:"required"`
	TotalAmount int        `json:"total_amount" validate:"required"`
	ExpiredAt   time.Time  `json:"expired_at" validate:"required"`
	PaidAt      *time.Time `json:"paid_at,omitempty"`
	Note        *string    `json:"note,omitempty"`
	TotalPrice  float64    `json:"total_price" validate:"true"`

	CreatedAt time.Time      `json:"created_at" `
	Items     []OrderItemDTO `json:"items" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;" validate:"required"`
}
