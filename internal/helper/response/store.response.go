package response

import (
	"payment-gateway/internal/model"
	"time"
)

type StoreResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	IsActive    bool    `json:"is_active"`
	OwnerID     string  `json:"owner_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ToStoreResponse(s model.Store) StoreResponse {
	return StoreResponse{
		ID:          s.ID,
		Name:        s.Name,
		Description: s.Description,
		IsActive:    s.IsActive,
		OwnerID:     s.OwnerID,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}

func ToStoreResponseList(stores []model.Store) []StoreResponse {
	responses := make([]StoreResponse, len(stores))
	for i, s := range stores {
		responses[i] = ToStoreResponse(s)
	}
	return responses
}
