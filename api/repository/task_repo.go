package repository

import (
	"cloud-native/api/models"
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	Create(task *models.Task) error
	UpdateTask(task *models.Task) error
	GetById(id string) (*models.Task, error)
	CompleteTask(id string) error
}

type taskRepository struct {
	taskCollection *mongo.Collection
	ctx            context.Context
}

func NewTaskRepository(mongo *mongo.Client) TaskRepository {
	taskCollection := mongo.Database("herlo").Collection("tasks")
	return &taskRepository{
		taskCollection: taskCollection,
		ctx:            context.TODO(),
	}
}

func (t taskRepository) Create(task *models.Task) error {
	_, err := t.taskCollection.InsertOne(t.ctx, task)
	if err != nil {
		return errors.Wrap(err, "taskRepo.Create")
	}
	return nil
}

func (t taskRepository) GetById(id string) (*models.Task, error) {
	taskFound := &models.Task{}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(err, "taskRepo.getById")
	}

	filter := bson.M{"_id": objectId}
	if err := t.taskCollection.FindOne(t.ctx, filter).Decode(taskFound); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, "taskRepo.GetByIf")
	}

	return taskFound, nil
}

func (t taskRepository) UpdateTask(task *models.Task) error {
	filter := bson.M{"_id": task.ID}
	update := bson.M{
		"$set": bson.M{
			"userId":      task.UserId,
			"description": task.Description,
			"title":       task.Title,
		},
	}
	_, err := t.taskCollection.UpdateOne(t.ctx, filter, update)
	if err != nil {
		return errors.Wrap(err, "taskRepo.CompleteTask")
	}
	return nil
}

func (t taskRepository) CompleteTask(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Wrap(err, "taskRepo.CompleteTask")
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{
		"$set": bson.M{
			"complete": true,
		}}

	_, err = t.taskCollection.UpdateOne(t.ctx, filter, update)

	if err != nil {
		return errors.Wrap(err, "taskRepo.CompleteTask")
	}

	return nil
}
