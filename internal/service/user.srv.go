package service

import (
	"context"
	"fmt"

	"payment-gateway/internal/helper"
	"payment-gateway/internal/model"
	"payment-gateway/internal/repository"

	"github.com/google/uuid"
)

type UserService interface {
	Create(ctx context.Context, m model.User) error
	GetMany(ctx context.Context, limit int, offset int) ([]model.User, int64, error)
	GetByRole(ctx context.Context, role model.Role, limit int, offset int) ([]model.User, int64, error)
	GetByID(ctx context.Context, ID string) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
}

type userService struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(ctx context.Context, m model.User) error {
	u := s.repo.User()
	m.ID = uuid.NewString()

	if !helper.IsEmailValid(m.Email) {
		return fmt.Errorf("email format invalid")
	}

	if err := u.
		Create(ctx, m); err != nil {
		return err
	}

	return nil
}

func (s *userService) GetMany(ctx context.Context, limit int, offset int) ([]model.User, int64, error) {
	u := s.repo.User()
	users, total, err := u.GetMany(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (s *userService) GetByRole(ctx context.Context, role model.Role, limit int, offset int) ([]model.User, int64, error) {
	u := s.repo.User()
	users, total, err := u.GetByRole(ctx, role, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (s *userService) GetByID(ctx context.Context, ID string) (*model.User, error) {
	u := s.repo.User()

	if err := uuid.Validate(ID); err != nil {
		return nil, err
	}

	user, err := u.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	u := s.repo.User()
	user, err := u.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
