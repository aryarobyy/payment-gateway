package response

import (
	"payment-gateway/internal/model"
	"time"
)

type OrderItemResponse struct {
	ID        string  `json:"id"`
	OrderID   string  `json:"order_id"`
	ProductID string  `json:"product_id"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	Subtotal  float64 `json:"subtotal"`
	CreatedAt time.Time `json:"created_at"`
}

func ToOrderItemResponse(item model.OrderItem) OrderItemResponse {
	return OrderItemResponse{
		ID:        item.ID,
		OrderID:   item.OrderID,
		ProductID: item.ProductID,
		Price:     item.Price,
		Quantity:  item.Quantity,
		Subtotal:  item.Subtotal,
		CreatedAt: item.CreatedAt,
	}
}

func ToOrderItemResponseList(items []model.OrderItem) []OrderItemResponse {
	responses := make([]OrderItemResponse, len(items))
	for i, item := range items {
		responses[i] = ToOrderItemResponse(item)
	}
	return responses
}

type OrderResponse struct {
	ID          string              `json:"id"`
	StoreID     string              `json:"store_id"`
	Status      model.Status        `json:"status"`
	TableID     string              `json:"table_id"`
	TotalAmount int                 `json:"total_amount"`
	TotalPrice  float64             `json:"total_price"`
	ExpiredAt   time.Time           `json:"expired_at"`
	PaidAt      *time.Time          `json:"paid_at,omitempty"`
	Note        *string             `json:"note,omitempty"`
	Items       []OrderItemResponse `json:"items"`
	CreatedAt   time.Time           `json:"created_at"`
}

func ToOrderResponse(o model.Order) OrderResponse {
	return OrderResponse{
		ID:          o.ID,
		StoreID:     o.StoreID,
		Status:      o.Status,
		TableID:     o.TableID,
		TotalAmount: o.TotalAmount,
		TotalPrice:  o.TotalPrice,
		ExpiredAt:   o.ExpiredAt,
		PaidAt:      o.PaidAt,
		Note:        o.Note,
		Items:       ToOrderItemResponseList(o.Items),
		CreatedAt:   o.CreatedAt,
	}
}

func ToOrderResponseList(orders []model.Order) []OrderResponse {
	responses := make([]OrderResponse, len(orders))
	for i, o := range orders {
		responses[i] = ToOrderResponse(o)
	}
	return responses
}
