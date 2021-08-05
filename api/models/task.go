package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	UserId      string             `json:"userId"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Complete    bool               `json:"complete"`
}

type CreateTaskRequest struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

func (r *CreateTaskRequest) Validate() *Error {
	return nil
}

type CreateTaskResponse struct {
	Id string `json:"id"`
}

type UpdateTaskRequest struct {
	Id          string `json:"id,omitempty"`
	UserId      string `json:"userId,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Complete    bool   `json:"complete,omitempty"`
}

func (r *UpdateTaskRequest) Validate() *Error {
	return nil
}
