package services

import (
	"fmt"
	"target/onboarding-assignment/models"
	"target/onboarding-assignment/repository/mocks"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserById(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	fmt.Println(mockUser)

	mockRepo := mocks.NewUserRepositoryInterface(t)
	mockRepo.On("GetUserById", mock.AnythingOfType("int")).Return(mockUser, nil)
	userService := UserService{Repo: mockRepo}
	mockUser, err = userService.GetUserById(2)
	fmt.Println(mockUser)
	fmt.Println(err)
	assert.NoError(t, err)
}
