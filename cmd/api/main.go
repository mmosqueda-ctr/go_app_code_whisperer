package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"go_app_for_code_whisperer/internal/handlers"
	"go_app_for_code_whisperer/pkg/database"
)

func main() {
	// Initialize the database
	// Replace with your actual connection string
	uri := "mongodb://localhost:27017"
	if err := database.InitDB(uri); err != nil {
		log.Fatalf("Could not initialize database: %s\n", err)
	}

	e := echo.New()
	handlers.RegisterRoutes(e)

	log.Println("Server starting on port 8080...")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
