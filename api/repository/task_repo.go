package repository

import (
	"new-proto/api/models"

	"github.com/pkg/errors"
)

type TaskRepository interface {
	Create(task models.Task) int
	UpdateTask(task models.Task) error
	GetById(id int) (models.Task, error)
	GetAll() []models.Task
	CompleteTask(id int) error
}

type taskRepository struct {
	idCounter      int
	taskCollection []models.Task
}

var ErrTaskNotFound = errors.New("task not found")

func NewTaskRepository() TaskRepository {
	return &taskRepository{}
}

func (t *taskRepository) Create(task models.Task) int {
	t.idCounter++
	task.ID = t.idCounter
	t.taskCollection = append(t.taskCollection, task)
	return t.idCounter
}

func (t *taskRepository) GetAll() []models.Task {
	return t.taskCollection
}

func (t *taskRepository) GetById(id int) (models.Task, error) {
	for _, v := range t.taskCollection {
		if v.ID == id {
			return v, nil
		}
	}
	return models.Task{}, ErrTaskNotFound
}

func (t *taskRepository) UpdateTask(task models.Task) error {
	for i, v := range t.taskCollection {
		if v.ID == task.ID {
			t.taskCollection[i] = task
			return nil
		}
	}
	return ErrTaskNotFound
}

func (t *taskRepository) CompleteTask(id int) error {
	for i, v := range t.taskCollection {
		if v.ID == id {
			t.taskCollection[i].Complete = true
			return nil
		}
	}
	return ErrTaskNotFound
}
