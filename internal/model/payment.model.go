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
