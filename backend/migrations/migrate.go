package migrations

import (
	"fmt"
	"log"
	"backend/models"
)

func RunMigrations() {
	// Apply any pending migrations
	fmt.Println("Running migrations...")

	// Auto-migrate the User model
	if err := models.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	fmt.Println("Migrations completed.")
}
