package services

import (
	"server/db"
	"server/dto"

	"gorm.io/gorm"
)

type ProductService struct {
	conn         *gorm.DB
	orderService OrderService
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{conn: db}
}

func (s *ProductService) GetProducts() []db.Product {
	var products []db.Product
	s.conn.Find(&products)
	return products
}

func (s *ProductService) GetProduct(id int) (db.Product, error) {
	var product db.Product
	if err := s.conn.Preload("Orders").Where("id = ?", id).First(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (s *ProductService) CreateProduct(productDto dto.CreateProductDto) db.Product {
	var product db.Product
	product, err := s.FindProductByName(productDto.Name)

	if err != nil {
		product = db.Product{Name: productDto.Name, Description: productDto.Description, Price: productDto.Price, Quantity: productDto.Quantity}
		s.conn.Create(&product)
	} else {
		product.Quantity += productDto.Quantity
		s.conn.Save(&product)
	}

	return product
}

func (s *ProductService) DeleteProduct(id int) {
	s.conn.Delete(&db.Product{}, id)
}

func (s *ProductService) CreateOrder(productID int, quantity int) (db.Order, error) {
	product, err := s.GetProduct(productID)

	if err != nil {
		return db.Order{}, err
	}
	order := s.orderService.CreateOrder(productID, quantity, product.Price)

	return order, nil
}

func (s *ProductService) FindProductByName(productName string) (db.Product, error) {
	var product db.Product
	if err := s.conn.Where("name = ?", productName).First(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (s *ProductService) GetOrdersForProduct(productID int) []db.Order {
	return s.orderService.GetOrdersForProduct(productID)
}
