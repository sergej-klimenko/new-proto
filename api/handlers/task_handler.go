package handlers

import (
	"net/http"
	"new-proto/api/models"
	"new-proto/api/services"
	"new-proto/api/utils"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type taskHandler struct {
	TaskSvc services.TaskService
}

func NewTaskHandler(taskSvc services.TaskService) http.Handler {
	handler := &taskHandler{
		TaskSvc: taskSvc,
	}

	r := chi.NewRouter()
	r.Post("/", handler.createTask)
	r.Get("/{id}", handler.getTask)
	r.Get("/", handler.getAllTasks)
	r.Put("/{id}", handler.updateTask)
	r.Post("/{id}/complete", handler.completeTask)
	return r
}

func (h taskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	task := &models.CreateTaskRequest{}
	if err := utils.DecodeAndValidate(r, task); err != nil {
		utils.WriteErrorResponse(w, err)
		return
	}

	newTaskId := h.TaskSvc.CreateTask(r.Context(), task)

	utils.WriteResponse(w, &models.CreateTaskResponse{Id: newTaskId}, 200)
}

func (h taskHandler) getAllTasks(w http.ResponseWriter, r *http.Request) {

	tasks := h.TaskSvc.GetAllTasks(r.Context())

	utils.WriteResponse(w, &tasks, 200)
}

func (h taskHandler) getTask(w http.ResponseWriter, r *http.Request) {
	id, err := getTaskId(r)

	if err != nil {
		utils.WriteErrorResponse(w, err)
		return
	}

	task, err := h.TaskSvc.GetTask(r.Context(), id)

	if err != nil {
		utils.WriteErrorResponse(w, err)
		return
	}

	utils.WriteResponse(w, task, 200)
}

func (h taskHandler) updateTask(w http.ResponseWriter, r *http.Request) {
	task := &models.UpdateTaskRequest{}
	if err := utils.DecodeAndValidate(r, task); err != nil {
		utils.WriteErrorResponse(w, err)
		return
	}

	err := h.TaskSvc.UpdateTask(r.Context(), task)

	if err != nil {
		utils.WriteErrorResponse(w, err)
		return
	}

	utils.WriteResponse(w, nil, http.StatusNoContent)
}

func (h taskHandler) completeTask(w http.ResponseWriter, r *http.Request) {
	id, err := getTaskId(r)

	if err != nil {
		utils.WriteErrorResponse(w, err)
		return
	}

	err = h.TaskSvc.CompleteTask(r.Context(), id)

	if err != nil {
		utils.WriteErrorResponse(w, err)
		return
	}

	utils.WriteResponse(w, nil, http.StatusNoContent)
}

func getTaskId(r *http.Request) (int, *models.Error) {
	taskId := chi.URLParam(r, "id")

	if taskId == "" {
		return 0, &models.Error{Code: http.StatusBadRequest, Message: "invalid task id"}
	}
	id, err := strconv.Atoi(taskId)

	if err != nil {
		return 0, &models.Error{Code: http.StatusBadRequest, Message: "invalid task id"}

	}
	return id, nil
}
