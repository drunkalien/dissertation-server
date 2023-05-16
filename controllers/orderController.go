package controllers

import (
	"net/http"
	"server/db"
	"server/dto"
	"server/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService *services.OrderService
}

func (controller *OrderController) GetOrders(c *gin.Context) {
	orders := controller.orderService.GetOrders()

	c.IndentedJSON(http.StatusOK, gin.H{"orders": orders})
}

func (controller *OrderController) CreateOrder(c *gin.Context) {
	var orderDto dto.CreateOrderDto
	if err := c.ShouldBindJSON(&orderDto); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := controller.orderService.CreateOrder(int(orderDto.ProductID), orderDto.Quantity, orderDto.Price)
	c.IndentedJSON(http.StatusCreated, gin.H{"order": order})
}

func (controller *OrderController) GetOrder(c *gin.Context) {
	strId := c.Param("id")

	id, err := strconv.Atoi(strId)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	order, err := controller.orderService.GetOrder(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"order": order})
}

func NewOrderController() *OrderController {
	return &OrderController{
		orderService: services.NewOrderService(db.DB),
	}
}
