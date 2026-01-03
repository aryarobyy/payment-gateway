package repository

import (
	"context"
	"fmt"

	"payment-gateway/internal/model"

	"gorm.io/gorm"
)

type ProductRepo interface {
	Create(ctx context.Context, m model.Product) error
	CreateBatch(ctx context.Context, m []model.Product) error
	GetMany(ctx context.Context, storeID string, limit int, offset int) ([]model.Product, int64, error)
	GetByCategory(ctx context.Context, category string, limit int, offset int) ([]model.Product, int64, error)
	GetByID(ctx context.Context, ID string) (*model.Product, error)
	GetByActive(ctx context.Context, isActive bool, limit int, offset int) ([]model.Product, int64, error)
	Update(ctx context.Context, m model.UpdateProduct, ID string) (*model.Product, error)
	GetPrice(ctx context.Context, ID string) (float64, error)
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return &productRepo{db: db}
}

func (r *productRepo) Create(ctx context.Context, m model.Product) error {
	return r.db.WithContext(ctx).
		Create(&m).
		Error
}

func (r *productRepo) CreateBatch(ctx context.Context, m []model.Product) error {
	return r.db.WithContext(ctx).
		CreateInBatches(&m, 100).
		Error
}

func (r *productRepo) GetMany(ctx context.Context, storeID string, limit int, offset int) ([]model.Product, int64, error) {
	var (
		total int64
		m     []model.Product
	)

	query := r.db.WithContext(ctx).
		Model([]model.Product{}).
		Where("store_id = ?", storeID)

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

func (r *productRepo) GetByID(ctx context.Context, ID string) (*model.Product, error) {
	m := model.Product{}

	if err := r.db.WithContext(ctx).
		First(&m, ID).
		Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *productRepo) GetByActive(ctx context.Context, isActive bool, limit int, offset int) ([]model.Product, int64, error) {
	var (
		total int64
		m     []model.Product
	)

	query := r.db.WithContext(ctx).
		Model([]model.Product{}).
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

func (r *productRepo) GetByCategory(ctx context.Context, category string, limit int, offset int) ([]model.Product, int64, error) {
	var (
		total int64
		m     []model.Product
	)

	query := r.db.WithContext(ctx).
		Model([]model.Product{}).
		Where("category = ?", category)

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

func (r *productRepo) Update(ctx context.Context, m model.UpdateProduct, id string) (*model.Product, error) {
	updateData := map[string]interface{}{}

	if m.Name != nil {
		updateData["name"] = *m.Name
	}
	if m.StoreID != nil {
		updateData["store_id"] = *m.StoreID
	}
	if m.Category != nil {
		updateData["category"] = *m.Category
	}
	if m.Price != nil {
		updateData["price"] = *m.Price
	}
	if m.IsActive != nil {
		updateData["is_active"] = *m.IsActive
	}

	if len(updateData) == 0 {
		return nil, fmt.Errorf("no field to update")
	}

	if err := r.db.WithContext(ctx).
		Model(&model.Product{}).
		Where("id = ?", id).
		Updates(updateData).
		Error; err != nil {
		return nil, err
	}

	updatedData := model.Product{}
	if err := r.db.WithContext(ctx).
		First(&updatedData, id).
		Error; err != nil {
		return nil, err
	}

	return &updatedData, nil
}

func (r *productRepo) GetPrice(ctx context.Context, ID string) (float64, error) {
	var price float64

	if err := r.db.WithContext(ctx).
		Where("id = ?", ID).
		Select("price").
		Error; err != nil {
		return 0.0, err
	}

	return price, nil
}
