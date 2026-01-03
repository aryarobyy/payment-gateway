package repository

import (
	"context"

	"payment-gateway/internal/model"

	"gorm.io/gorm"
)

type OrderItemRepo interface {
	Create(ctx context.Context, m model.OrderItem) error
	CreateBatch(ctx context.Context, m []model.OrderItem) error
	GetMany(ctx context.Context, orderID string) ([]model.OrderItem, error)
	GetByID(ctx context.Context, ID string) (*model.OrderItem, error)
}

type orderItemRepo struct {
	db *gorm.DB
}

func NewOrderItemRepo(db *gorm.DB) OrderItemRepo {
	return &orderItemRepo{db: db}
}

func (r *orderItemRepo) Create(ctx context.Context, m model.OrderItem) error {
	return r.db.WithContext(ctx).
		Create(&m).
		Error
}

func (r *orderItemRepo) CreateBatch(ctx context.Context, m []model.OrderItem) error {
	return r.db.WithContext(ctx).
		CreateInBatches(&m, 100).
		Error
}

func (r *orderItemRepo) GetMany(ctx context.Context, orderID string) ([]model.OrderItem, error) {
	m := []model.OrderItem{}

	query := r.db.WithContext(ctx).
		Model([]model.OrderItem{}).
		Where("order_id = ?", orderID)

	if err := query.
		Preload("Products").
		Find(&m).
		Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (r *orderItemRepo) GetByID(ctx context.Context, ID string) (*model.OrderItem, error) {
	m := model.OrderItem{}

	if err := r.db.WithContext(ctx).
		Preload("Products").
		First(&m, ID).
		Error; err != nil {
		return nil, err
	}
	return &m, nil
}

