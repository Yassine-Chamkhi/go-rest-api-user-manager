package services

import (
	"errors"
	"fmt"
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

func (svc *UserService) AddUser(user models.User) (returnedUser models.User, err error) {

	if user.Name == "" {
		err = errors.New("name property not specified")
		fmt.Println(err)
		return
	}

	if user.Age == 0 {
		err = errors.New("age property not specified")
		fmt.Println(err)
		return
	}

	return svc.Repo.AddUser(user)
}

func (svc *UserService) DeleteUserById(id int) error {
	return svc.Repo.DeleteUserById(id)
}
