package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"new-proto/api/models"
)

func WriteResponse(w http.ResponseWriter, res interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")

	content, err := json.Marshal(res)

	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)

	if _, err = w.Write(content); err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func WriteErrorResponse(w http.ResponseWriter, error *models.Error) {
	w.Header().Set("Content-Type", "application/json")

	if error.Code == http.StatusInternalServerError {
		fmt.Printf("%+v\n", error.Error)
	}

	content, err := json.Marshal(error)

	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(error.Code)

	if _, err = w.Write(content); err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
