package service

import (
	"context"
	"fmt"

	"payment-gateway/internal/helper"
	"payment-gateway/internal/model"
	"payment-gateway/internal/repository"

	"github.com/google/uuid"
)

type ProductService interface {
	Create(ctx context.Context, m model.PostProduct) error
	CreateBatch(ctx context.Context, m []model.Product) error
	GetMany(ctx context.Context, storeID string, limit int, offset int) ([]model.Product, int64, error)
	GetByCategory(ctx context.Context, category string, limit int, offset int) ([]model.Product, int64, error)
	GetByID(ctx context.Context, ID string) (*model.Product, error)
	GetByActive(ctx context.Context, isActive string, limit int, offset int) ([]model.Product, int64, error)
	Update(ctx context.Context, m model.UpdateProduct, id string) (*model.Product, error)
}

type productService struct {
	repo repository.Repository
}

func NewProductService(repo repository.Repository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) Create(ctx context.Context, m model.PostProduct) error {
	productRepo := s.repo.Product()
	m.ID = uuid.NewString()
	if m.Price <= 0 {
		return fmt.Errorf("invalid price for product %s", m.Name)
	}

	if err := helper.EmptyCheck(map[string]string{
		"store_id": m.StoreID,
		"name":     m.Name,
	}); err != nil {
		return err
	}

	product := model.Product{
		ID:       uuid.NewString(),
		StoreID:  m.StoreID,
		Name:     m.Name,
		Price:    m.Price,
		IsActive: m.IsActive,
		Category: m.Category,
	}

	if err := productRepo.Create(ctx, product); err != nil {
		return err
	}

	return nil
}

func (s *productService) CreateBatch(ctx context.Context, products []model.Product) error {
	repo := s.repo.Product()

	for i := range products {
		products[i].ID = uuid.NewString()

		if products[i].Price <= 0 {
			return fmt.Errorf("invalid price for product %s", products[i].Name)
		}

		if err := helper.EmptyCheck(map[string]string{
			"store_id": products[i].StoreID,
			"name":     products[i].Name,
		}); err != nil {
			return err
		}
	}

	return repo.CreateBatch(ctx, products)
}

func (s *productService) GetMany(ctx context.Context, storeID string, limit int, offset int) ([]model.Product, int64, error) {
	productRepo := s.repo.Product()
	products, total, err := productRepo.GetMany(ctx, storeID, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (s *productService) GetByCategory(ctx context.Context, category string, limit int, offset int) ([]model.Product, int64, error) {
	productRepo := s.repo.Product()
	products, total, err := productRepo.GetByCategory(ctx, category, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (s *productService) GetByID(ctx context.Context, ID string) (*model.Product, error) {
	productRepo := s.repo.Product()

	if err := uuid.Validate(ID); err != nil {
		return nil, err
	}

	product, err := productRepo.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) GetByActive(ctx context.Context, isActive string, limit int, offset int) ([]model.Product, int64, error) {
	productRepo := s.repo.Product()

	var isActiveBool bool
	switch isActive {
	case "true", "1", "yes":
		isActiveBool = true
	case "false", "0", "no":
		isActiveBool = false
	default:
		return nil, 0, fmt.Errorf("invalid isActive value: %s. Use true/false, 1/0, or yes/no", isActive)
	}

	products, total, err := productRepo.GetByActive(ctx, isActiveBool, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (s *productService) Update(ctx context.Context, m model.UpdateProduct, id string) (*model.Product, error) {
	productRepo := s.repo.Product()

	if err := uuid.Validate(id); err != nil {
		return nil, err
	}

	product, err := productRepo.Update(ctx, m, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
