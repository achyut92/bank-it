package dto

import (
	"bank-it/models"
	"strconv"
)

//Request DTO for account creation

// Validate using Gin's ShouldBindJson
type RequestAccount struct {
	AccountID int    `json:"account_id" binding:"required"`
	Balance   string `json:"initial_balance" binding:"required,numeric"`
}

type ResponseAccount struct {
	AccountID int    `json:"account_id"`
	Balance   string `json:"balance"`
}

// Convert DTO to DB Model
func (d *RequestAccount) ToModel() (*models.Account, error) {
	balanceFloat, err := strconv.ParseFloat(d.Balance, 64)
	if err != nil {
		return nil, err
	}
	return &models.Account{
		AccountID: d.AccountID,
		Balance:   balanceFloat,
	}, nil
}
