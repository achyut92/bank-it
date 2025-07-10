package db

import (
	"bank-it/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.Account{}, &models.Transaction{})
}
