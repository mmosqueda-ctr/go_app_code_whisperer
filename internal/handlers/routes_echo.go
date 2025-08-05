package handlers

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// User routes
	e.POST("/users", CreateUser)
	e.GET("/users/:id", GetUser)
	e.PUT("/users/:id", UpdateUser)
	e.DELETE("/users/:id", DeleteUser)

	// Product routes
	e.POST("/products", CreateProduct)
	e.GET("/products/:id", GetProduct)
	e.PUT("/products/:id", UpdateProduct)
	e.DELETE("/products/:id", DeleteProduct)

	// Order routes
	e.POST("/orders", CreateOrder)
	e.GET("/orders/:id", GetOrder)

	// Inventory routes
	e.PUT("/inventory", UpdateInventory)
}