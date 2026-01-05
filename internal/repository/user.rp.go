package repository

import (
	"context"

	"payment-gateway/internal/model"

	"gorm.io/gorm"
)

type UserRepo interface {
	Create(ctx context.Context, m model.User) error
	GetMany(ctx context.Context, limit int, offset int) ([]model.User, int64, error)
	GetByRole(ctx context.Context, role model.Role, limit int, offset int) ([]model.User, int64, error)
	GetByID(ctx context.Context, ID string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, m model.User) error {
	return r.db.WithContext(ctx).
		Create(&m).
		Error
}

func (r *userRepo) GetMany(ctx context.Context, limit int, offset int) ([]model.User, int64, error) {
	var (
		total int64
		m     []model.User
	)

	query := r.db.WithContext(ctx).
		Model([]model.User{})

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

func (r *userRepo) GetByID(ctx context.Context, ID string) (*model.User, error) {
	m := model.User{}

	if err := r.db.WithContext(ctx).
		First(&m, ID).
		Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *userRepo) GetByRole(ctx context.Context, role model.Role, limit int, offset int) ([]model.User, int64, error) {
	var (
		total int64
		m     []model.User
	)

	query := r.db.WithContext(ctx).
		Model([]model.User{}).
		Where("role = ?", role)

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

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	m := model.User{}
	query := r.db.WithContext(ctx).
		Model(model.User{}).
		Where("email = ?", email)

	if err := query.
		First(&m, email).
		Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *userRepo) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	m := model.User{}
	query := r.db.WithContext(ctx).
		Model(model.User{}).
		Where("username = ?", username)

	if err := query.
		First(&m, username).
		Error; err != nil {
		return nil, err
	}

	return &m, nil
}
