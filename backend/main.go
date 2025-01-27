package main

import (
	"fmt"
	"backend/models"
	"backend/routes"
	"log"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database connection
	models.InitDB()

	// Run database migration
	models.MigrateDB()

	// Set up routes
	routes.SetupRoutes()

	// Start the application
	fmt.Println("Server is running...")
}