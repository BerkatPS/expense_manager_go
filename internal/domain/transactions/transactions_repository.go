package transactions

import (
	"fmt"
	"github.com/BerkatPS/internal/domain/model"
	"gorm.io/gorm"
	"time"
)

type TransactionRepository interface {
	Create(transaction *model.Transaction) error
	FindAll() ([]*model.Transaction, error)
	CountTotalExpenseDaily(date time.Time) (float64, error)
	UpdateMonthlyReport(accountID int, amount float64, transactionType int) error
}

type transactionRepository struct {
	db *gorm.DB
}

func (r *transactionRepository) UpdateMonthlyReport(accountID int, amount float64, transactionType int) error {
	var reports model.MonthlyReport

	currentMonth := time.Now().Month()
	currentYear := time.Now().Year()

	err := r.db.Where("account_id = ? AND report_month = ? AND report_year = ?", accountID, currentMonth, currentYear).First(&reports).Error
	if err != nil {

		report := model.MonthlyReport{
			AccountID:   accountID,
			ReportMonth: int(currentMonth),
			ReportYear:  currentYear,
		}

		if transactionType == 1 {
			report.TotalIncome = amount
		} else if transactionType == 2 {
			report.TotalExpense = amount
		}

		err = r.db.Create(&report).Error
	}
	return fmt.Errorf("failed to update monthly report: %w", err)
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) CountTotalExpenseDaily(date time.Time) (float64, error) {
	var totalExpense float64
	err := r.db.Model(&model.Transaction{}).
		Where("transaction_date = ? AND DATE(created_at) = ?", "expense", date.Format("2006-01-02"), date).
		Select("SUM(amount)").
		Row().
		Scan(&totalExpense)
	return totalExpense, err
}

func (r *transactionRepository) Create(transaction *model.Transaction) error {
	return r.db.Create(&transaction).Error
}

func (r *transactionRepository) FindAll() ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	err := r.db.Find(&transactions).Error
	return transactions, err
}
