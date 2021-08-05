package middlewares

import (
	"cloud-native/api/models"
	"cloud-native/api/utils"
	"context"
	"net/http"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type currentUser string

var User = currentUser("currentUser")

func (c currentUser) String() string {
	return string(c)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := r.Header.Get("userId")

		if u == "" {
			apiError := &models.Error{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
			}
			utils.WriteErrorResponse(w, apiError)
			return
		}

		_, err := primitive.ObjectIDFromHex(u)
		if err != nil {
			apiError := &models.Error{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
				Error:   errors.Wrap(err, "middleware.AuthMiddleware"),
			}
			utils.WriteErrorResponse(w, apiError)
			return
		}

		ctx := context.WithValue(r.Context(), User, u)
		rr := r.WithContext(ctx)
		next.ServeHTTP(w, rr)
	})
}
