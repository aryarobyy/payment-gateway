package service

import (
	"context"

	"payment-gateway/internal/helper"
	"payment-gateway/internal/model"
	"payment-gateway/internal/repository"

	"github.com/google/uuid"
)

type OrderService interface {
	Create(ctx context.Context, m model.Order) error
	GetMany(ctx context.Context, limit int, offset int) ([]model.Order, int64, error)
	GetByStoreID(ctx context.Context, storeID string, limit int, offset int) ([]model.Order, int64, error)
	GetByStatus(ctx context.Context, status model.Status, limit int, offset int) ([]model.Order, int64, error)
	GetByID(ctx context.Context, ID string) (*model.Order, error)
	UpdateStatus(ctx context.Context, orderID string, status model.Status) error
}

type orderService struct {
	repo repository.Repository
}

func NewOrderService(
	repo repository.Repository,
) OrderService {
	return &orderService{
		repo: repo,
	}
}

func (s *orderService) Create(ctx context.Context, m model.Order) error {
	var totalPrice float64

	u := s.repo.Order()
	i := s.repo.OrderItem()
	m.ID = uuid.NewString()

	items, err := i.GetMany(ctx, m.ID)
	if err != nil {
		return err
	}

	for _, item := range items {
		totalPrice += item.Price
	}
	m.TotalPrice = totalPrice

	if err := helper.EmptyCheck(map[string]string{
		"store_id": m.StoreID,
		"table_id": m.TableID,
	}); err != nil {
		return err
	}

	if err := u.Create(ctx, m); err != nil {
		return err
	}

	return nil
}

func (s *orderService) GetMany(ctx context.Context, limit int, offset int) ([]model.Order, int64, error) {
	u := s.repo.Order()
	orders, total, err := u.GetMany(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}

func (s *orderService) GetByStoreID(ctx context.Context, storeID string, limit int, offset int) ([]model.Order, int64, error) {
	u := s.repo.Order()
	orders, total, err := u.GetByStoreID(ctx, storeID, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}

func (s *orderService) GetByStatus(ctx context.Context, status model.Status, limit int, offset int) ([]model.Order, int64, error) {
	u := s.repo.Order()
	orders, total, err := u.GetByStatus(ctx, status, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}

func (s *orderService) GetByID(ctx context.Context, ID string) (*model.Order, error) {
	u := s.repo.Order()

	if err := uuid.Validate(ID); err != nil {
		return nil, err
	}

	order, err := u.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *orderService) UpdateStatus(ctx context.Context, orderID string, status model.Status) error {
	u := s.repo.Order()

	if err := uuid.Validate(orderID); err != nil {
		return err
	}

	if err := u.UpdateStatus(ctx, orderID, status); err != nil {
		return err
	}

	return nil
}

