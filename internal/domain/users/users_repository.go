package users

import (
	"github.com/BerkatPS/internal/domain/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.Users) (*model.Users, error)
	FindAll() ([]*model.Users, error)
	Login(username string, password string) (*model.Users, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *model.Users) (*model.Users, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) FindAll() ([]*model.Users, error) {
	var users []*model.Users
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) Login(username string, password string) (*model.Users, error) {
	var user *model.Users
	err := r.db.Where("username = ? AND password = ?", username, password).First(&user).Error
	return user, err

}
