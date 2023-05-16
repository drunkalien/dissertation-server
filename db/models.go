package db

import "gorm.io/gorm"

const (
	RoleAdmin = iota
	RoleUser
)

type User struct {
	gorm.Model
	ID        uint   `json:"id"`
	Username  string `gorm:"unique" json:"username"`
	Password  string `json:"password"`
	Role      int    `json:"role"`
	CreatedAt int64  `gorm:"autoCreateTime" json:"created_at"`
	UpdateAt  int64  `gorm:"autoUpdatedTime" json:"updated_at"`
}

type Product struct {
	gorm.Model
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       int     `json:"price"`
	Quantity    int     `json:"quantity"`
	Orders      []Order `json:"orders"`
	CreatedAt   int64   `gorm:"autoCreateTime" json:"created_at"`
	UpdateAt    int64   `gorm:"autoUpdatedTime" json:"updated_at"`
}

type Order struct {
	gorm.Model
	ID         uint    `json:"id"`
	ProductID  uint    `json:"product_id"`
	Product    Product `gorm:"foreignKey:ID" json:"product"`
	Quantity   int     `json:"quantity"`
	TotalPrice int     `json:"total_price"`
	CreatedAt  int64   `gorm:"autoCreateTime" json:"created_at"`
	UpdateAt   int64   `gorm:"autoUpdatedTime" json:"updated_at"`
}
