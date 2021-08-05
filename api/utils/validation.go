package utils

import (
	"cloud-native/api/models"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type request interface {
	Validate() *models.Error
}

func DecodeAndValidate(r *http.Request, s request) *models.Error {
	if err := json.NewDecoder(r.Body).Decode(s); err != nil {
		return &models.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid request body",
			Error:   errors.Wrap(err, "json decoder request body"),
		}
	}

	if err := s.Validate(); err != nil {
		return err
	}

	return nil
}
