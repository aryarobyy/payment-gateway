package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"payment-gateway/internal/helper"
	"payment-gateway/internal/model"
	"payment-gateway/internal/repository"
)

type PaymentService interface {
	Create(ctx context.Context, m model.Payment, userID string) (*model.Payment, error)
	GetMany(ctx context.Context, storeID string, limit int, offset int) ([]model.Payment, int64, error)
	GetByID(ctx context.Context, id string) (*model.Payment, error)
	GetByOrderID(ctx context.Context, orderID string) (*model.Payment, error)
	GetByProviderRef(ctx context.Context, providerRef string) (*model.Payment, error)
	GetByStatus(ctx context.Context, status model.Status, limit int, offset int) ([]model.Payment, int64, error)
	UpdateStatus(ctx context.Context, paymentID string, status model.Status) error
	UpdateVerification(ctx context.Context, paymentID string, verifiedAt time.Time) error
}

type paymentService struct {
	repo repository.Repository
}

func NewPaymentService(repo repository.Repository) PaymentService {
	return &paymentService{repo: repo}
}

func (s *paymentService) Create(ctx context.Context, m model.Payment, userID string) (*model.Payment, error) {
	u := s.repo.Payment()

	if err := helper.EmptyCheck(map[string]string{
		"order_id": m.OrderID,
		"table_id": m.TableID,
		"provider": m.Provider,
	}); err != nil {
		return nil, err
	}

	if m.Amount <= 0 {
		return nil, fmt.Errorf("amount cannot be 0")
	}

	if m.Status == "" {
		m.Status = model.PENDING
	}

	if err := u.Create(ctx, m); err != nil {
		return nil, fmt.Errorf("failed to create payment: %w", err)
	}

	return &m, nil
}

func (s *paymentService) GetMany(ctx context.Context, storeID string, limit int, offset int) ([]model.Payment, int64, error) {
	u := s.repo.Payment()

	if storeID == "" {
		return nil, 0, fmt.Errorf("store id is empty")
	}

	payments, total, err := u.GetMany(ctx, storeID, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get payments by status: %w", err)
	}

	return payments, total, nil
}

func (s *paymentService) GetByID(ctx context.Context, id string) (*model.Payment, error) {
	u := s.repo.Payment()
	payment, err := u.GetByID(ctx, id)

	if id == "" {
		return nil, fmt.Errorf("id is empty")
	}

	if err != nil {
		return nil, fmt.Errorf("payment not found: %w", err)
	}

	return payment, nil
}

func (s *paymentService) GetByOrderID(ctx context.Context, orderID string) (*model.Payment, error) {
	u := s.repo.Payment()
	payment, err := u.GetByOrderID(ctx, orderID)

	if orderID == "" {
		return nil, fmt.Errorf("order id is empty")
	}

	if err != nil {
		return nil, fmt.Errorf("payment for order not found: %w", err)
	}

	return payment, nil
}

func (s *paymentService) GetByProviderRef(ctx context.Context, providerRef string) (*model.Payment, error) {
	u := s.repo.Payment()
	payment, err := u.GetByProviderRef(ctx, providerRef)

	if providerRef == "" {
		return nil, fmt.Errorf("store id is empty")
	}

	if err != nil {
		return nil, fmt.Errorf("payment with provider reference not found: %w", err)
	}

	return payment, nil
}

func (s *paymentService) GetByStatus(ctx context.Context, status model.Status, limit int, offset int) ([]model.Payment, int64, error) {
	u := s.repo.Payment()
	payments, total, err := u.GetByStatus(ctx, status, limit, offset)

	if status == "" {
		return nil, 0, fmt.Errorf("store id is empty")
	}

	if err != nil {
		return nil, 0, fmt.Errorf("failed to get payments by status: %w", err)
	}

	return payments, total, nil
}

func (s *paymentService) UpdateStatus(ctx context.Context, paymentID string, status model.Status) error {
	u := s.repo.Payment()

	if paymentID == "" {
		return fmt.Errorf("store id is empty")
	}

	if status != model.CREATED && status != model.EXPIRED && status != model.FAILED && status != model.PAID && status != model.PENDING {
		return fmt.Errorf("invalid status")
	}

	if err := u.UpdateStatus(ctx, paymentID, status); err != nil {
		return fmt.Errorf("failed to update payment status: %w", err)
	}

	return nil
}

func (s *paymentService) UpdateVerification(ctx context.Context, paymentID string, verifiedAt time.Time) error {
	u := s.repo.Payment()

	if paymentID == "" {
		return fmt.Errorf("payment id is empty")
	}

	if verifiedAt.IsZero() {
		return fmt.Errorf("verified at is empty")
	}

	payment, err := u.GetByID(ctx, paymentID)
	if err != nil {
		return fmt.Errorf("payment not found: %w", err)
	}

	if payment.Status != model.PENDING && payment.Status != model.PAID {
		return errors.New("only pending or paid payments can be verified")
	}

	if err := u.UpdateVerification(ctx, paymentID, verifiedAt); err != nil {
		return fmt.Errorf("failed to update payment verification: %w", err)
	}

	return nil
}
