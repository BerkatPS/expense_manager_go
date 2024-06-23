package accounts

import "github.com/BerkatPS/internal/domain/model"

type AccountService interface {
	CreateAccount(account *model.Accounts) (*model.Accounts, error)
	GetAllAccounts() ([]*model.Accounts, error)
}

type accountService struct {
	repository AccountRepository
}

func NewAccountService(repo AccountRepository) AccountService {
	return &accountService{repository: repo}
}

func (s *accountService) CreateAccount(account *model.Accounts) (*model.Accounts, error) {
	return s.repository.Create(account)
}

func (s *accountService) GetAllAccounts() ([]*model.Accounts, error) {
	return s.repository.FindAll()
}
