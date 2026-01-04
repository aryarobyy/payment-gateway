package repository

import (
	"context"
	"time"

	"payment-gateway/internal/model"

	"gorm.io/gorm"
)

type PaymentRepo interface {
	Create(ctx context.Context, m model.Payment) error
	GetMany(ctx context.Context, storeID string, limit int, offset int) ([]model.Payment, int64, error)
	GetByOrderID(ctx context.Context, orderID string) (*model.Payment, error)
	GetByProviderRef(ctx context.Context, providerRef string) (*model.Payment, error)
	GetByStatus(ctx context.Context, status model.Status, limit int, offset int) ([]model.Payment, int64, error)
	GetByID(ctx context.Context, ID string) (*model.Payment, error)
	UpdateStatus(ctx context.Context, paymentID string, status model.Status) error
	UpdateVerification(ctx context.Context, paymentID string, verifiedAt time.Time) error
}

type paymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) PaymentRepo {
	return &paymentRepo{db: db}
}

func (r *paymentRepo) Create(ctx context.Context, m model.Payment) error {
	return r.db.WithContext(ctx).
		Create(&m).
		Error
}

func (r *paymentRepo) GetMany(ctx context.Context, storeID string, limit int, offset int) ([]model.Payment, int64, error) {
	var (
		total int64
		m     []model.Payment
	)

	query := r.db.WithContext(ctx).
		Where("store_id = ?", storeID).
		Model([]model.Payment{})

	if err := query.
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}

	if err := query.
		Limit(limit).
		Offset(offset).
		Find(&m).
		Error; err != nil {
		return nil, 0, err
	}

	return m, total, nil
}

func (r *paymentRepo) GetByOrderID(ctx context.Context, orderID string) (*model.Payment, error) {
	m := model.Payment{}

	if err := r.db.WithContext(ctx).
		Where("order_id = ?", orderID).
		First(&m).
		Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *paymentRepo) GetByProviderRef(ctx context.Context, providerRef string) (*model.Payment, error) {
	m := model.Payment{}

	if err := r.db.WithContext(ctx).
		Where("provider_ref = ?", providerRef).
		First(&m).
		Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *paymentRepo) GetByStatus(ctx context.Context, status model.Status, limit int, offset int) ([]model.Payment, int64, error) {
	var (
		total int64
		m     []model.Payment
	)

	query := r.db.WithContext(ctx).
		Model([]model.Payment{}).
		Where("status = ?", status)

	if err := query.
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}

	if err := query.
		Limit(limit).
		Offset(offset).
		Find(&m).
		Error; err != nil {
		return nil, 0, err
	}

	return m, total, nil
}

func (r *paymentRepo) GetByID(ctx context.Context, ID string) (*model.Payment, error) {
	m := model.Payment{}

	if err := r.db.WithContext(ctx).
		First(&m, ID).
		Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *paymentRepo) UpdateStatus(ctx context.Context, paymentID string, status model.Status) error {
	return r.db.WithContext(ctx).
		Model(&model.Payment{}).
		Where("id = ?", paymentID).
		Update("status", status).
		Error
}

func (r *paymentRepo) UpdateVerification(ctx context.Context, paymentID string, verifiedAt time.Time) error {
	return r.db.WithContext(ctx).
		Model(&model.Payment{}).
		Where("id = ?", paymentID).
		Update("verified_at", verifiedAt).
		Error
}
