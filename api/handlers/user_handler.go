package handlers

import (
	"cloud-native/api/models"
	"cloud-native/api/services"
	"cloud-native/api/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type userHandler struct {
	UserSvc services.UserService
}

func NewUserHandler(userSvc services.UserService) http.Handler {
	handler := &userHandler{
		UserSvc: userSvc,
	}

	r := chi.NewRouter()
	r.Post("/", handler.createUser)
	return r
}

func (h userHandler) createUser(w http.ResponseWriter, r *http.Request) {
	user := &models.CreateUserRequest{}

	if err := utils.DecodeAndValidate(r, user); err != nil {
		utils.WriteErrorResponse(w, err)
		return
	}

	userId, err := h.UserSvc.CreateAccount(user)

	if err != nil {
		utils.WriteErrorResponse(w, err)
		return
	}

	utils.WriteResponse(w, &models.CreateUserResponse{Id: userId}, 200)
}
