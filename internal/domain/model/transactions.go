package model

import (
	"time"
)

type Transaction struct {
	TransactionID   int       `json:"transaction_id" gorm:"primaryKey;autoIncrement"`
	AccountID       int       `json:"account_id" gorm:"foreignKey:AccountID;references:AccountID"`
	CategoryID      int       `json:"category_id" gorm:"foreignKey:CategoryID;references:CategoryID"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount"`           // Changed to float64 to match decimal amount
	TransactionType int       `json:"transaction_type"` // Changed to string for type consistency
	CreatedAt       time.Time `json:"created_at"`
	//Account         Accounts  `gorm:"foreignKey:AccountID;references:AccountID"` // Many-to-One relationship with Account
	//Category        Categorys `gorm:"foreignKey:CategoryID;references:CategoryID"` // Many-to-One relationship with Category
}
