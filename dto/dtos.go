package dto

type CreateUserDto struct {
	Username string `json:"username"`
	Role     int    `json:"role"`
	Password string `json:"password"`
}

type CreateProductDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}

type SignInDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateOrderDto struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
	Price     int  `json:"price"`
}
