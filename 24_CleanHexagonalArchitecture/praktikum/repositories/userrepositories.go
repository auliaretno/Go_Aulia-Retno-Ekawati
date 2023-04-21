package repositories

import (
	"tugass/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsersRepository() ([]*models.User, error)
	CreateRepository(user models.User) (*models.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (u *userRepository) GetUsersRepository() ([]*models.User, error) {
	var users []*models.User

	if err := u.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) CreateRepository(user models.User) (*models.User, error) {
	if err := u.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}