package service

import (
	"payment-gateway/internal/repository"
)

type Service interface {
	Auth() authService
	User() userService
	Store() storeService
	Product() productService
	Payment() paymentService
	OrderItem() orderItemService
	Order() orderService
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Auth() authService       { return authService{repo: s.repo} }
func (s *service) User() userService       { return userService{repo: s.repo} }
func (s *service) Store() storeService     { return storeService{repo: s.repo} }
func (s *service) Product() productService { return productService{repo: s.repo} }

func (s *service) Payment() paymentService     { return paymentService{repo: s.repo} }
func (s *service) Order() orderService         { return orderService{repo: s.repo} }
func (s *service) OrderItem() orderItemService { return orderItemService{repo: s.repo} }
