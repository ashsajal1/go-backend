package initializers

import (
	// "go-api/initializers"
	"go-api/models"
)

func SyncDatabase() {
	// Migrate the schema
	DB.AutoMigrate(&models.User{})
}