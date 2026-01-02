package repository

import (
	"context"
	"fmt"

	"payment-gateway/internal/model"

	"gorm.io/gorm"
)

type OrderRepo interface {
	Create(ctx context.Context, m model.Order) error
	GetMany(ctx context.Context, limit int, offset int) ([]model.Order, int64, error)
	GetByStoreID(ctx context.Context, storeID string, limit int, offset int) ([]model.Order, int64, error)
	GetByStatus(ctx context.Context, status model.Status, limit int, offset int) ([]model.Order, int64, error)
	GetByID(ctx context.Context, ID string) (*model.Order, error)
	Update(ctx context.Context, m model.UpdateOrder, id string) (*model.Order, error)
	UpdateStatus(ctx context.Context, orderID string, status model.Status) error
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return &orderRepo{db: db}
}

func (r *orderRepo) Create(ctx context.Context, m model.Order) error {
	return r.db.WithContext(ctx).
		Create(&m).
		Error
}

func (r *orderRepo) GetMany(ctx context.Context, limit int, offset int) ([]model.Order, int64, error) {
	var (
		total int64
		m     []model.Order
	)

	query := r.db.WithContext(ctx).
		Model([]model.Order{})

	if err := query.
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}

	if err := query.
		Limit(limit).
		Offset(offset).
		Preload("Items").
		Find(&m).
		Error; err != nil {
		return nil, 0, err
	}

	return m, total, nil
}

func (r *orderRepo) GetByStoreID(ctx context.Context, storeID string, limit int, offset int) ([]model.Order, int64, error) {
	var (
		total int64
		m     []model.Order
	)

	query := r.db.WithContext(ctx).
		Model([]model.Order{}).
		Where("store_id = ?", storeID)

	if err := query.
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}

	if err := query.
		Limit(limit).
		Offset(offset).
		Preload("Items").
		Find(&m).
		Error; err != nil {
		return nil, 0, err
	}

	return m, total, nil
}

func (r *orderRepo) GetByStatus(ctx context.Context, status model.Status, limit int, offset int) ([]model.Order, int64, error) {
	var (
		total int64
		m     []model.Order
	)

	query := r.db.WithContext(ctx).
		Model([]model.Order{}).
		Where("status = ?", status)

	if err := query.
		Count(&total).
		Error; err != nil {
		return nil, 0, err
	}

	if err := query.
		Limit(limit).
		Offset(offset).
		Preload("Items").
		Find(&m).
		Error; err != nil {
		return nil, 0, err
	}

	return m, total, nil
}

func (r *orderRepo) GetByID(ctx context.Context, ID string) (*model.Order, error) {
	m := model.Order{}

	if err := r.db.WithContext(ctx).
		Preload("Items").
		First(&m, ID).
		Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *orderRepo) Update(ctx context.Context, m model.UpdateOrder, id string) (*model.Order, error) {
	updateData := map[string]interface{}{}

	if m.StoreID != nil {
		updateData["store_id"] = *m.StoreID
	}
	if m.Status != nil {
		updateData["status"] = *m.Status
	}
	if m.TableID != nil {
		updateData["table_id"] = *m.TableID
	}
	if m.TotalAmount != nil {
		updateData["total_amount"] = *m.TotalAmount
	}
	if m.ExpiredAt != nil {
		updateData["expired_at"] = *m.ExpiredAt
	}
	if m.PaidAt != nil {
		updateData["paid_at"] = *m.PaidAt
	}
	if m.Note != nil {
		updateData["note"] = *m.Note
	}
	if !m.UpdatedAt.IsZero() {
		updateData["updated_at"] = m.UpdatedAt
	}

	if len(updateData) == 0 {
		return nil, fmt.Errorf("no field to update")
	}

	if err := r.db.WithContext(ctx).
		Model(&model.Order{}).
		Where("id = ?", id).
		Updates(updateData).
		Error; err != nil {
		return nil, err
	}

	updatedData := model.Order{}
	if err := r.db.WithContext(ctx).
		Preload("Items").
		First(&updatedData, id).
		Error; err != nil {
		return nil, err
	}

	return &updatedData, nil
}

func (r *orderRepo) UpdateStatus(ctx context.Context, orderID string, status model.Status) error {
	return r.db.WithContext(ctx).
		Model(&model.Order{}).
		Where("id = ?", orderID).
		Update("status", status).
		Error
}