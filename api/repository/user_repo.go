package repository

import (
	"cloud-native/api/models"
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateAcount(user *models.User) error
	UsernameExists(username string) (bool, error)
	FindByUsername(username string) (*models.User, error)
	FindById(username string) (*models.User, error)
}

type userRepository struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserRepository(mongo *mongo.Client) UserRepository {
	userCollection := mongo.Database("golang_ecs").Collection("users")

	return &userRepository{
		userCollection: userCollection,
		ctx:            context.TODO(),
	}
}

func (u userRepository) CreateAcount(user *models.User) error {
	_, err := u.userCollection.InsertOne(u.ctx, user)

	if err != nil {
		return errors.Wrap(err, "ApiError inserting user")
	}

	return nil
}

func (u userRepository) UsernameExists(username string) (bool, error) {
	user, err := u.FindByUsername(username)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, errors.Wrap(err, "userRepo.UsernameExists")
	}

	if user != nil {
		return true, nil
	}

	return false, nil
}

func (u userRepository) FindByUsername(username string) (*models.User, error) {
	var userFound models.User

	filter := bson.D{primitive.E{Key: "username", Value: username}}

	if err := u.userCollection.FindOne(u.ctx, filter).Decode(&userFound); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, "userRepo.FindByUsername")
	}

	return &userFound, nil
}

func (u userRepository) FindById(userId string) (*models.User, error) {
	var userFound models.User

	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return nil, errors.Wrap(err, "userRepo.FindById")
	}

	filter := bson.M{"_id": objectId}

	if err := u.userCollection.FindOne(u.ctx, filter).Decode(&userFound); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, "userRepo.FindById")
	}

	return &userFound, nil
}
