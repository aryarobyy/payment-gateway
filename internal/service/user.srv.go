package service

import (
	"context"
	"fmt"
	"time"

	"payment-gateway/internal/helper"
	"payment-gateway/internal/model"
	"payment-gateway/internal/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, m model.RegisterCredential) error
	Login(ctx context.Context, m model.LoginCredential) (*model.User, string, error)
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

func (s *userService) Register(ctx context.Context, m model.RegisterCredential) error {
	u := s.repo.User()

	if err := helper.EmptyCheck(map[string]string{
		"email":    m.Email,
		"username": m.Username,
		"password": m.Password,
	}); err != nil {
		return err
	}

	if !helper.IsEmailValid(m.Email) {
		return fmt.Errorf("email format invalid")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(m.Password), 8)
	if err != nil {
		return err
	}

	user := model.User{
		ID:        uuid.NewString(),
		Username:  m.Username,
		Email:     m.Email,
		Password:  string(hashedPassword),
		Role:      model.Admin, // Sementara aja
		LastLogin: time.Now(),
	}

	if err := u.
		Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *userService) Login(ctx context.Context, m model.LoginCredential) (*model.User, string, error) {
	u := s.repo.User()
	if err := helper.EmptyCheck(map[string]string{
		"email":    m.Email,
		"password": m.Password,
	}); err != nil {
		return nil, "", err
	}
	if !helper.IsEmailValid(m.Email) {
		return nil, "", fmt.Errorf("invalid email format")
	}

	user, err := u.GetByEmail(ctx, m.Email)
	if err != nil {
		return nil, "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(user.Password)); err != nil {
		return nil, "", err
	}

	// token, err := helper.GenerateAccessToken(user)
	// if err != nil {
	// 	return nil, "", err
	// }

	return user, "", nil
}

func (s *userService) GetMany(ctx context.Context, limit int, offset int) ([]model.User, int64, error) {
	u := s.repo.User()
	users, total, err := u.GetMany(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (s *userService) GetByRole(ctx context.Context, role string, limit int, offset int) ([]model.User, int64, error) {
	u := s.repo.User()

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

	users, total, err := u.GetByRole(ctx, roleEnum, limit, offset)
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
