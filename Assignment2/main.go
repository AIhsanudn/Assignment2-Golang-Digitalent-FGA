package main

import (
	"Assignment2/controllers"
	"Assignment2/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Connect to the database
	database.Connect()

	// API routes
	v1 := router.Group("/api/v1")
	// order Endpoints
	{
		v1.POST("/orders", controllers.CreateOrder)
		v1.GET("/orders", controllers.GetOrders)
		v1.PUT("/orders/:id", controllers.UpdateOrder)
		v1.DELETE("/orders/:id", controllers.DeleteOrder)
	}

	// Item Endpoints
	{
		v1.POST("/items", controllers.CreateItem)
		v1.GET("/items", controllers.GetItems)
		v1.PUT("/items/:id", controllers.UpdateItem)
		v1.DELETE("/items/:id", controllers.DeleteItem)
	}

	// Run the server
	router.Run(":8000")
}
