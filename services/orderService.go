package services

import (
	"server/db"

	"gorm.io/gorm"
)

type OrderService struct {
	conn *gorm.DB
}

func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{conn: db}
}

func (s *OrderService) GetOrders() []db.Order {
	var orders []db.Order
	s.conn.Find(&orders)
	return orders
}

func (s *OrderService) GetOrder(id int) (db.Order, error) {
	var order db.Order
	if err := s.conn.Where("id = ?", id).First(&order).Error; err != nil {
		return order, err
	}

	return order, nil
}

func (s *OrderService) CreateOrder(productID int, quantity int, price int) db.Order {
	totalPrice := price * quantity
	order := db.Order{ProductID: uint(productID), Quantity: quantity, TotalPrice: totalPrice}
	s.conn.Create(&order)

	return order
}

func (s *OrderService) GetOrdersForProduct(productID int) []db.Order {
	var orders []db.Order
	s.conn.Where("product_id = ?", productID).Find(&orders)

	return orders
}
