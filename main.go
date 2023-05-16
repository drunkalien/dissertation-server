package main

import (
	"server/controllers"
	"server/db"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	db.Connect()

	r := gin.Default()

	r.Use(CORSMiddleware())

	ctrls := controllers.NewControllers()

	users := r.Group("/api/v1/users")
	{
		users.GET("/", ctrls.UserController.GetUsers)
		users.GET("/:id", ctrls.UserController.GetUser)
		users.GET("/me", ctrls.UserController.GetCurrentUser)
		users.POST("/create", ctrls.UserController.CreateUser)
		users.POST("/login", ctrls.UserController.SignIn)
		users.DELETE("/:id", ctrls.UserController.DeleteUser)
	}

	products := r.Group("/api/v1/products")
	{
		products.GET("/", ctrls.ProductController.GetProducts)
		products.GET("/:id", ctrls.ProductController.GetProduct)
		products.POST("/create", ctrls.ProductController.CreateProduct)
		products.GET("/:id/orders", ctrls.ProductController.GetOrdersForProduct)
	}

	orders := r.Group("/api/v1/orders")
	{
		orders.GET("/", ctrls.OrderController.GetOrders)
		orders.POST("/create", ctrls.OrderController.CreateOrder)
		orders.GET("/:id", ctrls.OrderController.GetOrder)
	}

	r.Run()
}
