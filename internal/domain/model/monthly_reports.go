package model

import (
	"time"
)

type MonthlyReport struct {
	ReportID     int       `json:"report_id" gorm:"primaryKey;autoIncrement"`
	AccountID    int       `json:"account_id" gorm:"foreignKey:AccountID;references:AccountID"`
	ReportMonth  int       `json:"report_month"`
	ReportYear   int       `json:"report_year"`
	TotalIncome  float64   `json:"total_income"`
	TotalExpense float64   `json:"total_expense"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	//Account      Accounts  `gorm:"foreignKey:AccountID;references:AccountID"` // Many-to-One relationship with Account
}
