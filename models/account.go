package models

type Account struct {
	AccountID int     `json:"account_id" gorm:"primaryKey"`
	Balance   float64 `json:"balance"`
}
