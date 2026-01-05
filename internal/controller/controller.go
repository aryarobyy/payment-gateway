package controller

import "payment-gateway/internal/service"

type Controller interface {
	User() UserController
	Product() ProductController
	Order() OrderController
	OrderItem() OrderItemController
	Payment() PaymentController
}

type controller struct {
	service service.Service
}

func NewContoller(service service.Service) Controller {
	return &controller{
		service: service,
	}
}

func (c *controller) User() UserController { return UserController{service: c.service} }

func (c *controller) Product() ProductController { return ProductController{service: c.service} }

func (c *controller) Order() OrderController { return OrderController{service: c.service} }

func (c *controller) OrderItem() OrderItemController { return OrderItemController{service: c.service} }

func (c *controller) Payment() PaymentController { return PaymentController{service: c.service} }
