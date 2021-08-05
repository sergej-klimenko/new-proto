package models

import (
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Username string             `json:"username"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
}

func (r *CreateUserRequest) Validate() *Error {
	var errors []string
	if r.Username == "" {
		errors = append(errors, "Username is required.")
	}

	if len(r.Username) < 3 {
		errors = append(errors, "Username must be at least 3 characters.")
	}

	if len(errors) == 0 {
		return nil
	}

	return &Error{
		Message: "Invalid request.",
		Code:    http.StatusBadRequest,
		Details: errors,
	}
}

type CreateUserResponse struct {
	Id string `json:"id"`
}
