package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"payment-gateway/internal/helper"
	"payment-gateway/internal/model"
	"payment-gateway/internal/repository"

	"github.com/google/uuid"
)

type PaymentService interface {
	Create(ctx context.Context, m model.PostPayment) (*model.Payment, error)
	GetMany(ctx context.Context, storeID string, limit int, offset int) ([]model.Payment, int64, error)
	GetByID(ctx context.Context, id string) (*model.Payment, error)
	GetByOrderID(ctx context.Context, orderID string) (*model.Payment, error)
	GetByProviderRef(ctx context.Context, providerRef string) (*model.Payment, error)
	GetByStatus(ctx context.Context, status string, limit int, offset int) ([]model.Payment, int64, error)
	UpdateStatus(ctx context.Context, paymentID string, status string) error
	UpdateVerification(ctx context.Context, paymentID string, verifiedAt time.Time) error
}

type paymentService struct {
	repo repository.Repository
}

func NewPaymentService(repo repository.Repository) PaymentService {
	return &paymentService{repo: repo}
}

func (s *paymentService) Create(ctx context.Context, m model.PostPayment) (*model.Payment, error) {
	paymentRepo := s.repo.Payment()

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

	payment := model.Payment{
		ID:          uuid.NewString(),
		OrderID:     m.OrderID,
		TableID:     m.TableID,
		Provider:    m.Provider,
		ProviderRef: m.ProviderRef,
		Amount:      m.Amount,
		Status:      m.Status,
		RawPayload:  m.RawPayload,
	}

	if err := paymentRepo.Create(ctx, payment); err != nil {
		return nil, fmt.Errorf("failed to create payment: %w", err)
	}

	return &payment, nil
}

func (s *paymentService) GetMany(ctx context.Context, storeID string, limit int, offset int) ([]model.Payment, int64, error) {
	paymentRepo := s.repo.Payment()

	if storeID == "" {
		return nil, 0, fmt.Errorf("store id is empty")
	}

	payments, total, err := paymentRepo.GetMany(ctx, storeID, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get payments by status: %w", err)
	}

	return payments, total, nil
}

func (s *paymentService) GetByID(ctx context.Context, id string) (*model.Payment, error) {
	paymentRepo := s.repo.Payment()
	payment, err := paymentRepo.GetByID(ctx, id)

	if id == "" {
		return nil, fmt.Errorf("id is empty")
	}

	if err != nil {
		return nil, fmt.Errorf("payment not found: %w", err)
	}

	return payment, nil
}

func (s *paymentService) GetByOrderID(ctx context.Context, orderID string) (*model.Payment, error) {
	paymentRepo := s.repo.Payment()
	payment, err := paymentRepo.GetByOrderID(ctx, orderID)

	if orderID == "" {
		return nil, fmt.Errorf("order id is empty")
	}

	if err != nil {
		return nil, fmt.Errorf("payment for order not found: %w", err)
	}

	return payment, nil
}

func (s *paymentService) GetByProviderRef(ctx context.Context, providerRef string) (*model.Payment, error) {
	paymentRepo := s.repo.Payment()
	payment, err := paymentRepo.GetByProviderRef(ctx, providerRef)

	if providerRef == "" {
		return nil, fmt.Errorf("store id is empty")
	}

	if err != nil {
		return nil, fmt.Errorf("payment with provider reference not found: %w", err)
	}

	return payment, nil
}

func (s *paymentService) GetByStatus(ctx context.Context, status string, limit int, offset int) ([]model.Payment, int64, error) {
	paymentRepo := s.repo.Payment()

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
		return nil, 0, fmt.Errorf("invalid status: %s", status)
	}

	payments, total, err := paymentRepo.GetByStatus(ctx, statusEnum, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get payments by status: %w", err)
	}

	return payments, total, nil
}

func (s *paymentService) UpdateStatus(ctx context.Context, paymentID string, status string) error {
	paymentRepo := s.repo.Payment()

	if paymentID == "" {
		return fmt.Errorf("payment id is empty")
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

	if err := paymentRepo.UpdateStatus(ctx, paymentID, statusEnum); err != nil {
		return fmt.Errorf("failed to update payment status: %w", err)
	}

	return nil
}

func (s *paymentService) UpdateVerification(ctx context.Context, paymentID string, verifiedAt time.Time) error {
	paymentRepo := s.repo.Payment()

	if paymentID == "" {
		return fmt.Errorf("payment id is empty")
	}

	if verifiedAt.IsZero() {
		return fmt.Errorf("verified at is empty")
	}

	payment, err := paymentRepo.GetByID(ctx, paymentID)
	if err != nil {
		return fmt.Errorf("payment not found: %w", err)
	}

	if payment.Status != model.PENDING && payment.Status != model.PAID {
		return errors.New("only pending or paid payments can be verified")
	}

	if err := paymentRepo.UpdateVerification(ctx, paymentID, verifiedAt); err != nil {
		return fmt.Errorf("failed to update payment verification: %w", err)
	}

	return nil
}
