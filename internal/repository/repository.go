package repository

import "gorm.io/gorm"

type Repository interface {
	User() userRepo
	Store() storeRepo
	Product() productRepo
	Payment() paymentRepo
	OrderItem() orderItemRepo
	Order() orderRepo
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) User() userRepo           { return userRepo{db: r.db} }
func (r *repository) Store() storeRepo         { return storeRepo{db: r.db} }
func (r *repository) Product() productRepo     { return productRepo{db: r.db} }
func (r *repository) Payment() paymentRepo     { return paymentRepo{db: r.db} }
func (r *repository) Order() orderRepo         { return orderRepo{db: r.db} }
func (r *repository) OrderItem() orderItemRepo { return orderItemRepo{db: r.db} }
