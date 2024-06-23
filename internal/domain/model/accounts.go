package model

import (
	"time"
)

type Accounts struct {
	AccountID   int       `json:"account_id" gorm:"primaryKey;autoIncrement"`
	UserID      int       `json:"user_id" gorm:"foreignKey:UserID;references:UserID"`
	AccountName string    `json:"account_name"`
	Balance     float64   `json:"balance"`
	CreatedAt   time.Time `json:"created_at"`
	//Transactions   []Transaction   `gorm:"foreignKey:AccountID;references:AccountID"` // One-to-Many relationship with Transaction
	//MonthlyReports []MonthlyReport `gorm:"foreignKey:AccountID;references:AccountID"` // One-to-Many relationship with MonthlyReport
}
