package repository

import (
	"context"

	"payment-gateway/internal/model"

	"gorm.io/gorm"
)

type AuthRepo interface {
	Create(ctx context.Context, m model.User) error
	VerifyEmail(ctx context.Context, e string) (*model.User, error)
}

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepo {
	return &authRepo{db: db}
}

func (r *authRepo) Create(ctx context.Context, m model.User) error {
	return r.db.WithContext(ctx).
		Create(&m).
		Error
}

func (r *authRepo) VerifyEmail(ctx context.Context, e string) (*model.User, error) {
	user := model.User{}
	if err := r.db.WithContext(ctx).
		Where("email = ?", e).
		First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
