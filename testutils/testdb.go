package testutils

import (
	"log"

	"bank-it/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewTestDB returns a fresh in-memory SQLite database
// with schema migrated, ready for isolated tests.
func NewTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open test database: %v", err)
	}

	err = db.AutoMigrate(&models.Account{}, &models.Transaction{})
	if err != nil {
		log.Fatalf("failed to migrate test database: %v", err)
	}

	return db
}
