package repository

import (
	"context"
	"fmt"

	"payment-gateway/internal/model"

	"gorm.io/gorm"
)

type StoreRepo interface {
	Create(ctx context.Context, m model.Store) error
	GetMany(ctx context.Context, limit int, offset int) ([]model.Store, int64, error)
	GetByOwnerID(ctx context.Context, ownerID string, limit int, offset int) ([]model.Store, int64, error)
	GetByIsActive(ctx context.Context, isActive bool, limit int, offset int) ([]model.Store, int64, error)
	GetByID(ctx context.Context, ID string) (*model.Store, error)
	Update(ctx context.Context, m model.UpdateStore, id string) (*model.Store, error)
}

type storeRepo struct {
	db *gorm.DB
}

func NewStoreRepo(db *gorm.DB) StoreRepo {
	return &storeRepo{db: db}
}

func (r *storeRepo) Create(ctx context.Context, m model.Store) error {
	return r.db.WithContext(ctx).
		Create(&m).
		Error
}

func (r *storeRepo) GetMany(ctx context.Context, limit int, offset int) ([]model.Store, int64, error) {
	var (
		total int64
		m     []model.Store
	)

	query := r.db.WithContext(ctx).
		Model([]model.Store{})

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

func (r *storeRepo) GetByOwnerID(ctx context.Context, ownerID string, limit int, offset int) ([]model.Store, int64, error) {
	var (
		total int64
		m     []model.Store
	)

	query := r.db.WithContext(ctx).
		Model([]model.Store{}).
		Where("owner_id = ?", ownerID)

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

func (r *storeRepo) GetByIsActive(ctx context.Context, isActive bool, limit int, offset int) ([]model.Store, int64, error) {
	var (
		total int64
		m     []model.Store
	)

	query := r.db.WithContext(ctx).
		Model([]model.Store{}).
		Where("is_active = ?", isActive)

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

func (r *storeRepo) GetByID(ctx context.Context, ID string) (*model.Store, error) {
	m := model.Store{}

	if err := r.db.WithContext(ctx).
		First(&m, ID).
		Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *storeRepo) Update(ctx context.Context, m model.UpdateStore, id string) (*model.Store, error) {
	updateData := map[string]interface{}{}

	if m.Name != nil {
		updateData["name"] = *m.Name
	}
	if m.Description != nil {
		updateData["description"] = *m.Description
	}
	if m.IsActive != nil {
		updateData["is_active"] = *m.IsActive
	}
	if m.OwnerID != nil {
		updateData["owner_id"] = *m.OwnerID
	}
	if !m.UpdatedAt.IsZero() {
		updateData["updated_at"] = m.UpdatedAt
	}

	if len(updateData) == 0 {
		return nil, fmt.Errorf("no field to update")
	}

	if err := r.db.WithContext(ctx).
		Model(&model.Store{}).
		Where("id = ?", id).
		Updates(updateData).
		Error; err != nil {
		return nil, err
	}

	updatedData := model.Store{}
	if err := r.db.WithContext(ctx).
		First(&updatedData, id).
		Error; err != nil {
		return nil, err
	}

	return &updatedData, nil
}