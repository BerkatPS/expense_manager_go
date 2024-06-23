package accounts

import (
	"github.com/BerkatPS/internal/domain/model"
	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(account *model.Accounts) (*model.Accounts, error)
	FindAll() ([]*model.Accounts, error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) Create(account *model.Accounts) (*model.Accounts, error) {
	err := r.db.Create(&account).Error
	return account, err
}

func (r *accountRepository) FindAll() ([]*model.Accounts, error) {
	var accounts []*model.Accounts
	err := r.db.Find(&accounts).Error
	return accounts, err
}
