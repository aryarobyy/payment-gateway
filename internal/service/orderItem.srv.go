package service

import (
	"context"
	"fmt"

	"payment-gateway/internal/model"
	"payment-gateway/internal/repository"

	"github.com/google/uuid"
)

type OrderItemService interface {
	Create(ctx context.Context, m model.OrderItem) error
	CreateBatch(ctx context.Context, m []model.OrderItem) error
	GetMany(ctx context.Context, orderID string) ([]model.OrderItem, error)
	GetByID(ctx context.Context, ID string) (*model.OrderItem, error)
	// GetPrices(ctx context.Context, ID []string) (float64, error)
}

type orderItemService struct {
	repo repository.Repository
}

func NewOrderItemService(
	repo repository.Repository,
) OrderItemService {
	return &orderItemService{
		repo: repo,
	}
}

func (s *orderItemService) Create(ctx context.Context, m model.OrderItem) error {
	u := s.repo.OrderItem()
	p := s.repo.Product()
	m.ID = uuid.NewString()

	if _, err := p.GetByID(ctx, m.ProductID); err != nil {
		return err
	}

	price, err := p.GetPrice(ctx, m.ProductID)
	if err != nil {
		return err
	}

	m.Price = price
	m.Subtotal = m.Price * float64(m.Quantity)

	if err := u.Create(ctx, m); err != nil {
		return err
	}

	return nil
}

func (s *orderItemService) CreateBatch(ctx context.Context, items []model.OrderItem) error {
	orderItemRepo := s.repo.OrderItem()
	productRepo := s.repo.Product()

	for i := range items {
		items[i].ID = uuid.NewString()
		if items[i].Quantity <= 0 {
			return fmt.Errorf("invalid quantity for product %s", items[i].ProductID)
		}

		product, err := productRepo.GetByID(ctx, items[i].ProductID)
		if err != nil {
			return err
		}

		items[i].ID = uuid.NewString()
		items[i].Price = product.Price
		items[i].Subtotal = product.Price * float64(items[i].Quantity)
	}

	if err := orderItemRepo.CreateBatch(ctx, items); err != nil {
		return err
	}

	return nil
}

func (s *orderItemService) GetMany(ctx context.Context, orderID string) ([]model.OrderItem, error) {
	u := s.repo.OrderItem()
	orderItems, err := u.GetMany(ctx, orderID)
	if err != nil {
		return nil, err
	}

	return orderItems, nil
}

func (s *orderItemService) GetByID(ctx context.Context, ID string) (*model.OrderItem, error) {
	u := s.repo.OrderItem()

	if err := uuid.Validate(ID); err != nil {
		return nil, err
	}

	orderItem, err := u.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return orderItem, nil
}

// func (s *orderItemService) GetPrices(ctx context.Context, ID []string) (float64, error) {
// 	var totalPrice float64
// 	for _, id := range ID {
// 		price, err := s.pRepo.GetPrice(ctx, id)
// 		if err != nil {
// 			return 0, fmt.Errorf("price unavailable")
// 		}
// 		totalPrice += price
// 	}
// 	return totalPrice, nil
// }
//

