package response

import (
	"payment-gateway/internal/model"
	"time"
)

type PaymentResponse struct {
	ID          string       `json:"id"`
	OrderID     string       `json:"order_id"`
	TableID     string       `json:"table_id"`
	Provider    string       `json:"provider"`
	ProviderRef string       `json:"provider_ref"`
	Amount      float64      `json:"amount"`
	Status      model.Status `json:"status"`
	VerifiedAt  *time.Time   `json:"verified_at,omitempty"`
	CreatedAt   time.Time    `json:"created_at"`
}

func ToPaymentResponse(p model.Payment) PaymentResponse {
	return PaymentResponse{
		ID:          p.ID,
		OrderID:     p.OrderID,
		TableID:     p.TableID,
		Provider:    p.Provider,
		ProviderRef: p.ProviderRef,
		Amount:      p.Amount,
		Status:      p.Status,
		VerifiedAt:  p.VerifiedAt,
		CreatedAt:   p.CreatedAt,
	}
}

func ToPaymentResponseList(payments []model.Payment) []PaymentResponse {
	responses := make([]PaymentResponse, len(payments))
	for i, p := range payments {
		responses[i] = ToPaymentResponse(p)
	}
	return responses
}
