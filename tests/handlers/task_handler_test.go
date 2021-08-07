package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"new-proto/api/handlers"
	"new-proto/api/services/mocks"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var taskSvcMock *mocks.TaskService
var taskHandler http.Handler

func init() {

	taskSvcMock = &mocks.TaskService{}
	taskHandler = handlers.NewTaskHandler(taskSvcMock)
}

func TestCreateTask_BadRequestBody(t *testing.T) {
	req := httptest.NewRequest("POST", "/", strings.NewReader(""))
	rec := httptest.NewRecorder()

	taskHandler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "invalid request body")

}
