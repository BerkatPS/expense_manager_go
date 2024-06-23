package users

import "github.com/BerkatPS/internal/domain/model"

type UserService interface {
	CreateUser(user *model.Users) (*model.Users, error)
	GetAllUsers() ([]*model.Users, error)
	Login(username string, password string) (*model.Users, error)
}

type userService struct {
	repository UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repository: repo}
}

func (s *userService) CreateUser(user *model.Users) (*model.Users, error) {
	return s.repository.Create(user)
}

func (s *userService) GetAllUsers() ([]*model.Users, error) {
	return s.repository.FindAll()
}

func (s *userService) Login(username string, password string) (*model.Users, error) {
	return s.repository.Login(username, password)
}