package dto

import (
	"bank-it/models"
	"strconv"
)

//Request DTO to transfer money

type Transfer struct {
	SourceAccountID      int    `json:"source_account_id" binding:"required"`
	DestinationAccountID int    `json:"destination_account_id" binding:"required"`
	Amount               string `json:"amount" binding:"required,numeric"`
}

func (d *Transfer) ToModel() (*models.Transaction, error) {
	amountFloat, err := strconv.ParseFloat(d.Amount, 64)
	if err != nil {
		return nil, err
	}
	return &models.Transaction{
		SourceAccountID:      d.SourceAccountID,
		DestinationAccountID: d.DestinationAccountID,
		Amount:               amountFloat,
	}, nil
}
