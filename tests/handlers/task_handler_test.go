package handlers_test

import (
	"cloud-native/api/handlers"
	"cloud-native/api/services/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/stretchr/testify/assert"
)

var taskSvcMock *mocks.TaskService
var taskHandler http.Handler

func init() {

	taskSvcMock = &mocks.TaskService{}
	taskHandler = handlers.NewTaskHandler(taskSvcMock)
}

func TestCreateTask_BadRequestBody(t *testing.T) {
	id := primitive.NewObjectID()
	req := httptest.NewRequest("POST", "/", strings.NewReader(""))
	req.Header.Set("userId", id.Hex())
	rec := httptest.NewRecorder()

	taskHandler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "invalid request body")

}

func TestCreateTask_MissingUserIdHeader(t *testing.T) {
	req := httptest.NewRequest("POST", "/", strings.NewReader(""))
	rec := httptest.NewRecorder()

	taskHandler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.Contains(t, rec.Body.String(), "Unauthorized")

}
