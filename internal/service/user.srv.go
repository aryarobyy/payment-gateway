package service

import (
	"context"
	"fmt"

	"payment-gateway/internal/model"
	"payment-gateway/internal/repository"

	"github.com/google/uuid"
)

type UserService interface {
	GetMany(ctx context.Context, limit int, offset int) ([]model.User, int64, error)
	GetByRole(ctx context.Context, role string, limit int, offset int) ([]model.User, int64, error)
	GetByID(ctx context.Context, ID string) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
}

type userService struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetMany(ctx context.Context, limit int, offset int) ([]model.User, int64, error) {
	userRepo := s.repo.User()
	users, total, err := userRepo.GetMany(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (s *userService) GetByRole(ctx context.Context, role string, limit int, offset int) ([]model.User, int64, error) {
	userRepo := s.repo.User()

	var roleEnum model.Role
	switch role {
	case "admin":
		roleEnum = model.Admin
	case "super_admin":
		roleEnum = model.SuperAdmin
	case "owner":
		roleEnum = model.Owner
	case "staff":
		roleEnum = model.Staff
	default:
		return nil, 0, fmt.Errorf("invalid role: %s", role)
	}

	users, total, err := userRepo.GetByRole(ctx, roleEnum, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (s *userService) GetByID(ctx context.Context, ID string) (*model.User, error) {
	userRepo := s.repo.User()

	if err := uuid.Validate(ID); err != nil {
		return nil, err
	}

	user, err := userRepo.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	userRepo := s.repo.User()
	user, err := userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
