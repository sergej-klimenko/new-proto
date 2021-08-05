package services

import (
	"cloud-native/api/models"
	"cloud-native/api/repository"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	CreateAccount(user *models.CreateUserRequest) (string, *models.Error)
	GetUserById(userId string) (*models.User, *models.Error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return userService{
		userRepo: userRepo,
	}
}

func (u userService) CreateAccount(request *models.CreateUserRequest) (string, *models.Error) {
	userExists, err := u.userRepo.UsernameExists(request.Username)

	if err != nil {
		return "", &models.Error{
			Error:   err,
			Message: "unable to create account",
			Code:    http.StatusInternalServerError,
		}
	}

	if userExists {
		return "", &models.Error{
			Message: "username already exists",
			Code:    http.StatusBadRequest,
		}
	}

	user := &models.User{
		Id:       primitive.NewObjectID(),
		Username: strings.TrimSpace(request.Username),
	}

	err = u.userRepo.CreateAcount(user)
	if err != nil {
		return "", &models.Error{
			Error:   err,
			Message: "unable to create account",
			Code:    http.StatusInternalServerError,
		}
	}

	return user.Id.Hex(), nil
}

func (u userService) GetUserById(userId string) (*models.User, *models.Error) {
	user, err := u.userRepo.FindById(userId)

	if err != nil {
		return nil, &models.Error{
			Error:   err,
			Message: "unable to find account",
			Code:    http.StatusNotFound,
		}
	}

	if user == nil {
		return nil, &models.Error{
			Message: "account does not exist",
			Code:    http.StatusNotFound,
		}
	}

	return user, nil
}
