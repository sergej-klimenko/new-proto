package services

import (
	"context"
	"net/http"
	"new-proto/api/models"
	"new-proto/api/repository"
	"strings"
)

type TaskService interface {
	CreateTask(ctx context.Context, task *models.CreateTaskRequest) int
	UpdateTask(ctx context.Context, task *models.UpdateTaskRequest) *models.Error
	GetTask(ctx context.Context, id int) (*models.Task, *models.Error)
	GetAllTasks(ctx context.Context) []models.Task
	CompleteTask(ctx context.Context, id int) *models.Error
}

type taskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(taskRepo repository.TaskRepository) TaskService {
	return &taskService{
		taskRepo: taskRepo,
	}
}

func (t *taskService) CreateTask(ctx context.Context, task *models.CreateTaskRequest) int {
	newTask := models.Task{
		Description: strings.TrimSpace(task.Description),
		Title:       strings.TrimSpace(task.Title),
		Complete:    false,
	}
	id := t.taskRepo.Create(newTask)
	return id
}

func (t *taskService) GetAllTasks(ctx context.Context) []models.Task {
	return t.taskRepo.GetAll()
}

func (t *taskService) GetTask(ctx context.Context, id int) (*models.Task, *models.Error) {

	task, err := t.taskRepo.GetById(id)

	if err != nil {
		return nil, &models.Error{
			Code:    http.StatusNotFound,
			Message: "task does not exist",
			Error:   err,
		}
	}

	return &task, nil
}

func (t *taskService) UpdateTask(ctx context.Context, task *models.UpdateTaskRequest) *models.Error {

	taskFound, err := t.taskRepo.GetById(task.Id)
	if err != nil {
		return &models.Error{
			Code:    http.StatusNotFound,
			Message: "task does not exist",
			Error:   err,
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

func (t *taskService) CompleteTask(ctx context.Context, id int) *models.Error {

	task, err := t.taskRepo.GetById(id)

	if err != nil {
		return &models.Error{
			Code:    http.StatusNotFound,
			Message: "task does not exist",
			Error:   err,
		}
	}

	if err := t.taskRepo.CompleteTask(task.ID); err != nil {
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
