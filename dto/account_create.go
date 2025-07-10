package dto

import (
	"bank-it/models"
	"strconv"
)

//Request DTO for account creation

// Validate using Gin's ShouldBindJson
type CreateAccount struct {
	AccountID int    `json:"account_id" binding:"required"`
	Balance   string `json:"initial_balance" binding:"required,numeric"`
}

// Convert DTO to DB Model
func (d *CreateAccount) ToModel() (*models.Account, error) {
	balanceFloat, err := strconv.ParseFloat(d.Balance, 64)
	if err != nil {
		return nil, err
	}
	return &models.Account{
		AccountID: d.AccountID,
		Balance:   balanceFloat,
	}, nil
}
