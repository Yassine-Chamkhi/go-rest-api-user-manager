package services

import (
	"target/onboarding-assignment/models"
	"target/onboarding-assignment/repository"
)

type UserServiceInterface interface {
	GetUserById(id int) (models.User, error)
	GetAllUsers() ([]models.User, error)
	AddUser(user models.User) (models.User, error)
	DeleteUserById(id int) error
}

type UserService struct {
	Repo repository.UserRepositoryInterface
}

func (svc *UserService) GetUserById(id int) (models.User, error) {
	return svc.Repo.GetUserById(id)
}

func (svc *UserService) GetAllUsers() ([]models.User, error) {
	return svc.Repo.GetAllUsers()
}

func (svc *UserService) AddUser(user models.User) (models.User, error) {
	return svc.Repo.AddUser(user)
}

func (svc *UserService) DeleteUserById(id int) error {
	return svc.Repo.DeleteUserById(id)
}
