package services

import (
	"errors"
	"go-rest-api/models"
	"go-rest-api/repository/mocks"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserById(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockRepo := mocks.NewUserRepositoryInterface(t)

	mockRepo.On("GetUserById", 5).Return(mockUser, nil)
	mockRepo.On("GetUserById", 0).Return(models.User{}, errors.New("sql: no rows in result set"))

	userService := UserService{Repo: mockRepo}

	//Testing case where user id is a valid id; example 5
	gotUser, err := userService.GetUserById(5)
	assert.NotEmpty(t, gotUser)
	assert.NoError(t, err)

	//Testing case where user id is not a valid id; example 0
	gotUser, err = userService.GetUserById(0)
	assert.Empty(t, gotUser)
	assert.EqualError(t, err, "sql: no rows in result set")
}

func TestGetAllUsers(t *testing.T) {
	var mockUsers []models.User
	err := faker.FakeData(&mockUsers)
	assert.NoError(t, err)

	mockRepo := mocks.NewUserRepositoryInterface(t)

	mockRepo.On("GetAllUsers").Return(mockUsers, nil)

	userService := UserService{Repo: mockRepo}

	gotUsers, err := userService.GetAllUsers()
	assert.NotEmpty(t, gotUsers)
	assert.NoError(t, err)
	assert.Len(t, gotUsers, len(mockUsers))
}

func TestAddUser(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockRepo := mocks.NewUserRepositoryInterface(t)
	mockRepo.On("AddUser", mock.AnythingOfType("models.User")).Return(mockUser, nil)

	userService := UserService{Repo: mockRepo}

	//Testing case where user is a valid user
	returnedUser, err := userService.AddUser(mockUser)
	assert.NotEmpty(t, returnedUser)
	assert.NoError(t, err)

	//Testing case where user name is empty
	returnedUser, err = userService.AddUser(models.User{Age: 25})
	assert.Empty(t, returnedUser)
	assert.EqualError(t, err, "name property not specified")

	//Testing case where user age is empty
	returnedUser, err = userService.AddUser(models.User{Name: "testName"})
	assert.Empty(t, returnedUser)
	assert.EqualError(t, err, "age property not specified")

}

func TestDeleteUserById(t *testing.T) {

	mockRepo := mocks.NewUserRepositoryInterface(t)
	mockRepo.On("DeleteUserById", 5).Return(nil)
	mockRepo.On("DeleteUserById", 0).Return(errors.New("sql: no rows in result set"))

	userService := UserService{Repo: mockRepo}

	//Testing case where user id is a valid id; example 5
	err := userService.DeleteUserById(5)
	assert.NoError(t, err)

	//Testing case where user id is not a valid id; example 0
	err = userService.DeleteUserById(0)
	assert.Error(t, err)
	assert.EqualError(t, err, "sql: no rows in result set")

}
