package handlers

import (
	"net/http"
	"strconv"

	"bank-it/dto"
	"bank-it/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//Handler to manage Account related operations

func RegisterAccountRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/accounts", createAccountHandler(db))
	r.GET("/accounts/:account_id", getAccountHandler(db))
}

func createAccountHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dtoAccount dto.RequestAccount

		//Validate and bind request body
		if err := c.ShouldBindJSON(&dtoAccount); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		modelAccount, err := dtoAccount.ToModel()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid balance format"})
			return
		}
		//Create Account + Handles duplicate account creation
		if err := db.Create(&modelAccount).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create account"})
			return
		}
		c.Status(http.StatusOK)
	}
}

func getAccountHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		accountID, err := strconv.Atoi(c.Param("account_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
			return
		}

		var acc models.Account
		if err := db.First(&acc, "account_id = ?", accountID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}

		response := dto.ResponseAccount{
			AccountID: acc.AccountID,
			Balance:   strconv.FormatFloat(acc.Balance, 'f', -2, 64),
		}

		c.JSON(http.StatusOK, response)
	}
}
