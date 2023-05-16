package controllers

type Controllers struct {
	UserController    *UserController
	ProductController *ProductController
	OrderController   *OrderController
}

func NewControllers() *Controllers {
	return &Controllers{
		UserController:    NewUserController(),
		ProductController: NewProductController(),
		OrderController:   NewOrderController(),
	}
}
