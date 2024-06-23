package transactions

import (
	"fmt"
	"github.com/BerkatPS/internal/domain/model"
	"time"
)

type TransactionService interface {
	CreateTransaction(accountID int, categoryID int, transactionType int, amount float64) error
	GetAllTransactions() ([]*model.Transaction, error)
	CountTotalExpenseDaily(date time.Time) (float64, error)
}

type transactionService struct {
	repository TransactionRepository
}

func NewTransactionService(repo TransactionRepository) TransactionService {
	return &transactionService{repository: repo}
}

func (s *transactionService) CountTotalExpenseDaily(date time.Time) (float64, error) {
	return s.repository.CountTotalExpenseDaily(date)

}
func (s *transactionService) CreateTransaction(accountID int, categoryID int, transactionType int, amount float64) error {
	transaction := &model.Transaction{
		AccountID:       accountID,
		CategoryID:      categoryID,
		Amount:          amount,
		TransactionType: transactionType,
		CreatedAt:       time.Time{},
	}

	err := s.repository.Create(transaction)
	if err != nil {
		return err
	}

	err = s.repository.UpdateMonthlyReport(accountID, amount, transactionType)
	if err != nil {
		return err
	}

	return fmt.Errorf("failed to create transaction: %w", err)
}

func (s *transactionService) GetAllTransactions() ([]*model.Transaction, error) {
	return s.repository.FindAll()
}
