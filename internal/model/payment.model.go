package model

import "time"

type Payment struct {
	ID          string     `json:"id" gorm:"type:uuid;primaryKey"`
	OrderID     string     `json:"order_id"`
	TableID     string     `json:"table_id" validate:"required"`
	Provider    string     `json:"provider"`
	ProviderRef string     `json:"provider_ref"`
	Amount      float64    `json:"amount"`
	Status      Status     `json:"status"`
	RawPayload  string     `json:"raw_payload"`
	VerifiedAt  *time.Time `json:"verified_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}

type UpdatePayment struct {
	OrderID     *string  `json:"order_id,omitempty"`
	TableID     *string  `json:"table_id,omitempty" validate:"required"`
	Provider    *string  `json:"provider,omitempty"`
	ProviderRef *string  `json:"provider_ref,omitempty"`
	Amount      *float64 `json:"amount,omitempty"`
	Status      *Status  `json:"status,omitempty"`
	RawPayload  *string  `json:"raw_payload,omitempty"`
	VerifiedAt  *time.Time `json:"verified_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at"`
}
