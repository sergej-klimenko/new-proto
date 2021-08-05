package services

import (
	"cloud-native/api/middlewares"
	"cloud-native/api/models"
	"cloud-native/api/repository"
	"context"
	"errors"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskService interface {
	CreateTask(ctx context.Context, task *models.CreateTaskRequest) (string, *models.Error)
	UpdateTask(ctx context.Context, task *models.UpdateTaskRequest) *models.Error
	GetTask(ctx context.Context, id string) (*models.Task, *models.Error)
	CompleteTask(ctx context.Context, id string) *models.Error
}

type taskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(taskRepo repository.TaskRepository) TaskService {
	return taskService{
		taskRepo: taskRepo,
	}
}

func (t taskService) CreateTask(ctx context.Context, task *models.CreateTaskRequest) (string, *models.Error) {
	activeUser, err := GetActiveUser(ctx)

	if err != nil {
		return "", &models.Error{
			Code:    http.StatusUnauthorized,
			Message: "unauthorized",
		}
	}

	newTask := &models.Task{
		ID:          primitive.NewObjectID(),
		UserId:      activeUser,
		Description: strings.TrimSpace(task.Description),
		Title:       strings.TrimSpace(task.Title),
		Complete:    false,
	}

	if err := t.taskRepo.Create(newTask); err != nil {
		return "", &models.Error{
			Code:    http.StatusInternalServerError,
			Message: "could not create task",
			Error:   err,
		}
	}

	return newTask.ID.Hex(), nil
}

func (t taskService) GetTask(ctx context.Context, id string) (*models.Task, *models.Error) {
	activeUser, err := GetActiveUser(ctx)
	if err != nil {
		return nil, &models.Error{
			Code:    http.StatusUnauthorized,
			Message: "unauthorized",
		}
	}

	task, err := t.taskRepo.GetById(id)

	if err != nil {
		return nil, &models.Error{
			Code:    http.StatusNotFound,
			Message: "task does not exist",
			Error:   err,
		}
	}

	if task.UserId != activeUser {
		return nil, &models.Error{
			Code:    http.StatusUnauthorized,
			Message: "unauthorized",
		}
	}

	return task, nil
}

func (t taskService) UpdateTask(ctx context.Context, task *models.UpdateTaskRequest) *models.Error {
	activeUser, err := GetActiveUser(ctx)

	if err != nil {
		return &models.Error{
			Code:    http.StatusUnauthorized,
			Message: "unauthorized",
		}
	}

	taskFound, err := t.taskRepo.GetById(task.Id)
	if err != nil {
		return &models.Error{
			Code:    http.StatusNotFound,
			Message: "task does not exist",
			Error:   err,
		}
	}

	if taskFound.UserId != activeUser {
		return &models.Error{
			Code:    http.StatusUnauthorized,
			Message: "unauthorized",
		}
	}

	taskFound.Title = task.Title
	taskFound.Description = task.Description

	if err := t.taskRepo.UpdateTask(taskFound); err != nil {
		return &models.Error{
			Code:    http.StatusInternalServerError,
			Message: "could not complete task",
			Error:   err,
		}
	}
	return nil
}

func (t taskService) CompleteTask(ctx context.Context, id string) *models.Error {
	activeUser, err := GetActiveUser(ctx)

	if err != nil {
		return &models.Error{
			Code:    http.StatusUnauthorized,
			Message: "unauthorized",
		}
	}

	task, err := t.taskRepo.GetById(id)

	if err != nil {
		return &models.Error{
			Code:    http.StatusNotFound,
			Message: "task does not exist",
			Error:   err,
		}
	}

	if task.UserId != activeUser {
		return &models.Error{
			Code:    http.StatusUnauthorized,
			Message: "unauthorized",
		}
	}

	if err := t.taskRepo.CompleteTask(id); err != nil {
		return &models.Error{
			Code:    http.StatusInternalServerError,
			Message: "could not complete task",
			Error:   err,
		}
	}

	// send email/text message
	// place a message on a queue for another service to process
	// etc

	return nil
}

func GetActiveUser(ctx context.Context) (string, error) {
	activeUser := ctx.Value(middlewares.User).(string)
	if activeUser == "" {
		return "", errors.New("Unauthorized")
	}
	return activeUser, nil
}
