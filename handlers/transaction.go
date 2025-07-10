package handlers

import (
	"net/http"

	"bank-it/dto"
	"bank-it/enums"
	"bank-it/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//Handler to manage Transaction related operations

func RegisterTransactionRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/transactions", createTransactionHandler(db))
}

func createTransactionHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dtoTxn dto.Transfer
		if err := c.ShouldBindJSON(&dtoTxn); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		if dtoTxn.DestinationAccountID == dtoTxn.SourceAccountID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Source and Destination ID cannot be same"})
			return
		}

		txn, convertError := dtoTxn.ToModel()
		if convertError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount format"})
			return
		}

		var responseCode int = http.StatusOK
		var responseBody any = nil
		referenceId := uuid.New().String()

		err := db.Transaction(func(tx *gorm.DB) error {
			var src models.Account
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&src, "account_id = ?", txn.SourceAccountID).Error; err != nil {
				responseCode = http.StatusNotFound
				responseBody = gin.H{"error": "Source account not found"}
				return err
			}

			if src.Balance < txn.Amount {
				responseCode = http.StatusBadRequest
				responseBody = gin.H{"error": "Insufficient funds"}
				return gorm.ErrInvalidTransaction
			}

			var dst models.Account
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&dst, "account_id = ?", txn.DestinationAccountID).Error; err != nil {
				responseCode = http.StatusNotFound
				responseBody = gin.H{"error": "Destination account not found"}
				return err
			}

			src.Balance -= txn.Amount

			dst.Balance += txn.Amount

			if err := tx.Save(&src).Error; err != nil {
				return err
			}
			if err := tx.Save(&dst).Error; err != nil {
				return err
			}

			debitTxn := models.Transaction{
				SourceAccountID:      txn.SourceAccountID,
				DestinationAccountID: txn.DestinationAccountID,
				Amount:               txn.Amount,
				Balance:              src.Balance,
				TransactionType:      enums.Debit,
				ReferenceId:          referenceId,
			}

			creditTxn := models.Transaction{
				SourceAccountID:      txn.SourceAccountID,
				DestinationAccountID: txn.DestinationAccountID,
				Amount:               txn.Amount,
				Balance:              dst.Balance,
				TransactionType:      enums.Credit,
				ReferenceId:          referenceId,
			}

			if err := tx.Create(&debitTxn).Error; err != nil {
				return err
			}
			if err := tx.Create(&creditTxn).Error; err != nil {
				return err
			}

			return nil
		})

		if responseBody != nil {
			c.JSON(responseCode, responseBody)
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction failed"})
			return
		}

		c.Status(http.StatusOK)
	}
}
