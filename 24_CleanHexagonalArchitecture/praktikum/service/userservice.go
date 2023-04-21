package services

import (
	"tugass/models"
	"tugass/repositories"
)

type UserService interface {
	GetUsersService() ([]*models.User, error)
	CreateService(user models.User) (*models.User, error)
}

type userService struct {
	userR repositories.UserRepository
}

func NewUserService(userR repositories.UserRepository) UserService {
	return &userService{
		userR: userR,
	}
}

func (u *userService) GetUsersService() ([]*models.User, error) {
	users, err := u.userR.GetUsersRepository()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userService) CreateService(user models.User) (*models.User, error) {
	userR, err := u.userR.CreateRepository(user)
	if err != nil {
		return nil, err
	}

	return userR, nil
}