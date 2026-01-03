package service

import (
	"context"

	"payment-gateway/internal/model"
	"payment-gateway/internal/repository"

	"github.com/google/uuid"
)

type StoreService interface {
	Create(ctx context.Context, m model.Store) error
	GetMany(ctx context.Context, limit int, offset int) ([]model.Store, int64, error)
	GetByOwnerID(ctx context.Context, ownerID string, limit int, offset int) ([]model.Store, int64, error)
	GetByIsActive(ctx context.Context, isActive bool, limit int, offset int) ([]model.Store, int64, error)
	GetByID(ctx context.Context, ID string) (*model.Store, error)
	Update(ctx context.Context, m model.UpdateStore, id string) (*model.Store, error)
}

type storeService struct {
	repo repository.Repository
}

func NewStoreService(repo repository.Repository) StoreService {
	return &storeService{repo: repo}
}

func (s *storeService) Create(ctx context.Context, m model.Store) error {
	u := s.repo.Store()
	uR := s.repo.User()
	m.ID = uuid.NewString()

	if _, err := uR.GetByID(ctx, m.OwnerID); err != nil {
		return err
	}

	if err := u.Create(ctx, m); err != nil {
		return err
	}

	return nil
}

func (s *storeService) GetMany(ctx context.Context, limit int, offset int) ([]model.Store, int64, error) {
	u := s.repo.Store()
	stores, total, err := u.GetMany(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return stores, total, nil
}

func (s *storeService) GetByOwnerID(ctx context.Context, ownerID string, limit int, offset int) ([]model.Store, int64, error) {
	u := s.repo.Store()
	stores, total, err := u.GetByOwnerID(ctx, ownerID, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return stores, total, nil
}

func (s *storeService) GetByIsActive(ctx context.Context, isActive bool, limit int, offset int) ([]model.Store, int64, error) {
	u := s.repo.Store()
	stores, total, err := u.GetByIsActive(ctx, isActive, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return stores, total, nil
}

func (s *storeService) GetByID(ctx context.Context, ID string) (*model.Store, error) {
	u := s.repo.Store()

	if err := uuid.Validate(ID); err != nil {
		return nil, err
	}

	store, err := u.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (s *storeService) Update(ctx context.Context, m model.UpdateStore, id string) (*model.Store, error) {
	u := s.repo.Store()

	if err := uuid.Validate(id); err != nil {
		return nil, err
	}

	store, err := u.Update(ctx, m, id)
	if err != nil {
		return nil, err
	}

	return store, nil
}
