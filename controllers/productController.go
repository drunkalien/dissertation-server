package controllers

import (
	"net/http"
	"server/db"
	"server/dto"
	"server/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService *services.ProductService
}

func (controller *ProductController) GetProducts(c *gin.Context) {
	products := controller.productService.GetProducts()

	c.IndentedJSON(http.StatusOK, gin.H{"products": products})
}

func (controller *ProductController) CreateProduct(c *gin.Context) {
	var productDto dto.CreateProductDto
	err := c.ShouldBindJSON(&productDto)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := controller.productService.CreateProduct(productDto)

	c.IndentedJSON(http.StatusCreated, gin.H{"product": product})
}

func (controller *ProductController) GetProduct(c *gin.Context) {
	strId := c.Param("id")

	id, err := strconv.Atoi(strId)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	user, err := controller.productService.GetProduct(id)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"product": user})
}

func (controller *ProductController) GetOrdersForProduct(c *gin.Context) {
	strId := c.Param("id")

	id, err := strconv.Atoi(strId)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	orders := controller.productService.GetOrdersForProduct(id)

	c.IndentedJSON(http.StatusOK, gin.H{"orders": orders})
}

func NewProductController() *ProductController {
	productService := services.NewProductService(db.DB)
	return &ProductController{productService: productService}
}
