package repository

import (
	"context"
	"fmt"

	"payment-gateway/internal/model"

	"gorm.io/gorm"
)

type OrderItemRepo interface {
	Create(ctx context.Context, m model.OrderItem) error
	CreateBatch(ctx context.Context, m []model.OrderItem) error
	GetMany(ctx context.Context, orderID string, limit int, offset int) ([]model.OrderItem, int64, error)
	GetByOrderID(ctx context.Context, orderID string) ([]model.OrderItem, error)
	GetByID(ctx context.Context, ID string) (*model.OrderItem, error)
	Update(ctx context.Context, m model.UpdateOrderItem, id string) (*model.OrderItem, error)
	DeleteByOrderID(ctx context.Context, orderID string) error
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

func (r *orderItemRepo) GetMany(ctx context.Context, orderID string, limit int, offset int) ([]model.OrderItem, int64, error) {
	var (
		total int64
		m     []model.OrderItem
	)

	query := r.db.WithContext(ctx).
		Model([]model.OrderItem{}).
		Where("order_id = ?", orderID)

	if err := query.
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}

	if err := query.
		Limit(limit).
		Offset(offset).
		Preload("Products").
		Find(&m).
		Error; err != nil {
		return nil, 0, err
	}

	return m, total, nil
}

func (r *orderItemRepo) GetByOrderID(ctx context.Context, orderID string) ([]model.OrderItem, error) {
	var m []model.OrderItem

	if err := r.db.WithContext(ctx).
		Preload("Products").
		Where("order_id = ?", orderID).
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

func (r *orderItemRepo) Update(ctx context.Context, m model.UpdateOrderItem, id string) (*model.OrderItem, error) {
	updateData := map[string]interface{}{}

	if m.OrderID != nil {
		updateData["order_id"] = *m.OrderID
	}
	if m.ProductID != nil {
		updateData["product_id"] = *m.ProductID
	}
	if m.Price != nil {
		updateData["price"] = *m.Price
	}
	if m.Quantity != nil {
		updateData["quantity"] = *m.Quantity
	}
	if m.Subtotal != nil {
		updateData["subtotal"] = *m.Subtotal
	}
	if m.UpdatedAt != nil {
		updateData["updated_at"] = *m.UpdatedAt
	}

	if len(updateData) == 0 {
		return nil, fmt.Errorf("no field to update")
	}

	if err := r.db.WithContext(ctx).
		Model(&model.OrderItem{}).
		Where("id = ?", id).
		Updates(updateData).
		Error; err != nil {
		return nil, err
	}

	updatedData := model.OrderItem{}
	if err := r.db.WithContext(ctx).
		Preload("Products").
		First(&updatedData, id).
		Error; err != nil {
		return nil, err
	}

	return &updatedData, nil
}

func (r *orderItemRepo) DeleteByOrderID(ctx context.Context, orderID string) error {
	return r.db.WithContext(ctx).
		Where("order_id = ?", orderID).
		Delete(&model.OrderItem{}).
		Error
}