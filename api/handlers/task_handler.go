package handlers

import (
	"cloud-native/api/middlewares"
	"cloud-native/api/models"
	"cloud-native/api/services"
	"cloud-native/api/utils"
	"net/http"

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
	r.Use(middlewares.AuthMiddleware)
	r.Post("/", handler.createTask)
	r.Get("/{id}", handler.getTask)
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

	newTaskId, err := h.TaskSvc.CreateTask(r.Context(), task)

	if err != nil {
		utils.WriteErrorResponse(w, err)
		return
	}

	utils.WriteResponse(w, &models.CreateTaskResponse{Id: newTaskId}, 200)
}

func (h taskHandler) getTask(w http.ResponseWriter, r *http.Request) {
	taskId := chi.URLParam(r, "id")

	if taskId == "" {
		err := &models.Error{Code: http.StatusBadRequest, Message: "invalid taskId"}
		utils.WriteErrorResponse(w, err)
		return
	}

	task, err := h.TaskSvc.GetTask(r.Context(), taskId)

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
	taskId := chi.URLParam(r, "id")

	if taskId == "" {
		err := &models.Error{Code: http.StatusBadRequest, Message: "invalid taskId"}
		utils.WriteErrorResponse(w, err)
		return
	}

	err := h.TaskSvc.CompleteTask(r.Context(), taskId)

	if err != nil {
		utils.WriteErrorResponse(w, err)
		return
	}

	utils.WriteResponse(w, nil, http.StatusNoContent)
}
