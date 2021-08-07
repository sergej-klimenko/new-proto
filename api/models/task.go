package models

type CreateTaskRequest struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

func (r *CreateTaskRequest) Validate() *Error {
	if r.Title == "" || r.Description == "" {
		return &Error{}
	}
	return nil
}

type CreateTaskResponse struct {
	Id int `json:"id"`
}

type UpdateTaskRequest struct {
	Id          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Complete    bool   `json:"complete,omitempty"`
}

func (r *UpdateTaskRequest) Validate() *Error {
	return nil
}
