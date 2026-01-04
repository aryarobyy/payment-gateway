package service

import (
	"payment-gateway/internal/repository"
)

type Service interface {
	User() userService
	Store() storeService
	Product() productService
	// Payment() payment
	OrderItem() orderItemService
	Order() orderService
}

type service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) User() userService       { return userService{repo: *s.repo} }
func (s *service) Store() storeService     { return storeService{repo: *s.repo} }
func (s *service) Product() productService { return productService{repo: *s.repo} }

// func  (s *service) Payment() payment     { return paymentRepo{db: r.db} }
func (s *service) Order() orderService         { return orderService{repo: *s.repo} }
func (s *service) OrderItem() orderItemService { return orderItemService{repo: *s.repo} }
