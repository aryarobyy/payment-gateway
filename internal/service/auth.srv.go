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

type AuthService interface {
	Register(ctx context.Context, m model.RegisterCredential) error
	Login(ctx context.Context, m model.LoginCredential) (*model.User, string, error)
	RefreshToken(ctx context.Context, rToken string) (string, error)
}

type authService struct {
	repo repository.Repository
}

func NewAuthService(repo repository.Repository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Register(ctx context.Context, m model.RegisterCredential) error {
	authRepo := s.repo.Auth()

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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := model.User{
		ID:        uuid.NewString(),
		Username:  m.Username,
		Email:     m.Email,
		Password:  string(hashedPassword),
		Role:      model.Staff,
		LastLogin: time.Now(),
	}

	if err := authRepo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *authService) Login(ctx context.Context, m model.LoginCredential) (*model.User, string, error) {
	authRepo := s.repo.Auth()

	if err := helper.EmptyCheck(map[string]string{
		"email":    m.Email,
		"password": m.Password,
	}); err != nil {
		return nil, "", err
	}

	if !helper.IsEmailValid(m.Email) {
		return nil, "", fmt.Errorf("invalid email format")
	}

	user, err := authRepo.VerifyEmail(ctx, m.Email)
	if err != nil {
		return nil, "", err
	}

	context.WithValue(ctx, "userID", user.ID)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(m.Password))
	if err != nil {
		return nil, "", fmt.Errorf("invalid password")
	}

	token, err := helper.GenerateAccessToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *authService) RefreshToken(ctx context.Context, rToken string) (string, error) {
	userSrv := s.repo.User()

	claims, err := helper.ParseAndValidateToken(rToken)
	if err != nil {
		return "", err
	}

	user, err := userSrv.GetByID(ctx, claims.ID)
	if err != nil {
		return "", err
	}

	newToken, err := helper.GenerateAccessToken(user)
	if err != nil {
		return "", err
	}

	return newToken, nil
}
