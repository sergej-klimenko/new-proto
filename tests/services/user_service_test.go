package services_test

import (
	"cloud-native/api/models"
	"cloud-native/api/repository/mocks"
	"cloud-native/api/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/pkg/errors"
)

func TestUserService_UsernameExistFails(t *testing.T) {
	userRepoMock := &mocks.UserRepository{}
	userSvc := services.NewUserService(userRepoMock)

	newUserRequest := &models.CreateUserRequest{
		Username: "brian",
	}

	expectedError := errors.New("error")

	userRepoMock.On("UsernameExists", mock.Anything).Return(false, expectedError, nil)

	_, err := userSvc.CreateAccount(newUserRequest)

	userRepoMock.AssertExpectations(t)
	assert.True(t, errors.Is(expectedError, err.Error))
	assert.Equal(t, 500, err.Code)
	assert.Contains(t, err.Message, "unable to create account")
}

func TestUserService_UsernameTaken(t *testing.T) {
	userRepoMock := &mocks.UserRepository{}
	userSvc := services.NewUserService(userRepoMock)

	newUserRequest := &models.CreateUserRequest{
		Username: "brian",
	}

	userRepoMock.On("UsernameExists", mock.Anything).Return(true, nil, nil)

	_, err := userSvc.CreateAccount(newUserRequest)

	userRepoMock.AssertExpectations(t)
	assert.Equal(t, 400, err.Code)
	assert.Contains(t, err.Message, "username already exists")
}
