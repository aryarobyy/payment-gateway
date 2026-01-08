package service

import (
	"context"
	"fmt"

	"payment-gateway/internal/helper"
	"payment-gateway/internal/model"
	"payment-gateway/internal/repository"

	"github.com/google/uuid"
)

type OrderService interface {
	Create(ctx context.Context, m model.Order) error
	GetMany(ctx context.Context, limit int, offset int) ([]model.Order, int64, error)
	GetByStoreID(ctx context.Context, storeID string, limit int, offset int) ([]model.Order, int64, error)
	GetByStatus(ctx context.Context, status string, limit int, offset int) ([]model.Order, int64, error)
	GetByID(ctx context.Context, ID string) (*model.Order, error)
	UpdateStatus(ctx context.Context, orderID string, status string) error
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

	orderRepo := s.repo.Order()
	itemRepo := s.repo.OrderItem()
	m.ID = uuid.NewString()

	items, err := itemRepo.GetMany(ctx, m.ID)
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

	if err := orderRepo.Create(ctx, m); err != nil {
		return err
	}

	return nil
}

func (s *orderService) GetMany(ctx context.Context, limit int, offset int) ([]model.Order, int64, error) {
	orderRepo := s.repo.Order()
	orders, total, err := orderRepo.GetMany(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}

func (s *orderService) GetByStoreID(ctx context.Context, storeID string, limit int, offset int) ([]model.Order, int64, error) {
	orderRepo := s.repo.Order()
	orders, total, err := orderRepo.GetByStoreID(ctx, storeID, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}

func (s *orderService) GetByStatus(ctx context.Context, status string, limit int, offset int) ([]model.Order, int64, error) {
	orderRepo := s.repo.Order()

	var statusEnum model.Status
	switch status {
	case "created":
		statusEnum = model.CREATED
	case "pending":
		statusEnum = model.PENDING
	case "paid":
		statusEnum = model.PAID
	case "expired":
		statusEnum = model.EXPIRED
	case "failed":
		statusEnum = model.FAILED
	default:
		return nil, 0, fmt.Errorf("invalid status: %s", status)
	}

	orders, total, err := orderRepo.GetByStatus(ctx, statusEnum, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}

func (s *orderService) GetByID(ctx context.Context, ID string) (*model.Order, error) {
	orderRepo := s.repo.Order()

	if err := uuid.Validate(ID); err != nil {
		return nil, err
	}

	order, err := orderRepo.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *orderService) UpdateStatus(ctx context.Context, orderID string, status string) error {
	orderRepo := s.repo.Order()

	if err := uuid.Validate(orderID); err != nil {
		return err
	}

	// Convert status string to enum
	var statusEnum model.Status
	switch status {
	case "created":
		statusEnum = model.CREATED
	case "pending":
		statusEnum = model.PENDING
	case "paid":
		statusEnum = model.PAID
	case "expired":
		statusEnum = model.EXPIRED
	case "failed":
		statusEnum = model.FAILED
	default:
		return fmt.Errorf("invalid status: %s", status)
	}

	if err := orderRepo.UpdateStatus(ctx, orderID, statusEnum); err != nil {
		return err
	}

	return nil
}
