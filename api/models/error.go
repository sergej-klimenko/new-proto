package models

type Error struct {
	Message string      `json:"message"`
	Error   error       `json:"-"`
	Code    int         `json:"code"`
	Details interface{} `json:"details,omitempty"`
}
