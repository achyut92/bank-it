package dto

import (
	"bank-it/models"
	"strconv"
)

//Request DTO for account creation

// Validate using Gin's ShouldBindJson
type Account struct {
	AccountID int    `json:"account_id" binding:"required"`
	Balance   string `json:"initial_balance" binding:"required,numeric"`
}

// Convert DTO to DB Model
func (d *Account) ToModel() (*models.Account, error) {
	balanceFloat, err := strconv.ParseFloat(d.Balance, 64)
	if err != nil {
		return nil, err
	}
	return &models.Account{
		AccountID: d.AccountID,
		Balance:   balanceFloat,
	}, nil
}
