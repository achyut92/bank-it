package models

import "bank-it/enums"

type Transaction struct {
	SourceAccountID      int     `json:"source_account_id"`
	DestinationAccountID int     `json:"destination_account_id"`
	Amount               float64 `json:"amount"`
	Balance              float64
	TransactionType      enums.TransactionType
	ReferenceId          string
}
