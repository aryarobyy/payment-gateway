package response

import (
	"payment-gateway/internal/model"
	"time"
)

type ProductResponse struct {
	ID        string  `json:"id"`
	StoreID   string  `json:"store_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	IsActive  bool    `json:"is_active"`
	Category  string  `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToProductResponse(p model.Product) ProductResponse {
	return ProductResponse{
		ID:        p.ID,
		StoreID:   p.StoreID,
		Name:      p.Name,
		Price:     p.Price,
		IsActive:  p.IsActive,
		Category:  p.Category,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func ToProductResponseList(products []model.Product) []ProductResponse {
	responses := make([]ProductResponse, len(products))
	for i, p := range products {
		responses[i] = ToProductResponse(p)
	}
	return responses
}
