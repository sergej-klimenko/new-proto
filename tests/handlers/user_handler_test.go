package handlers_test

import (
	"cloud-native/api/handlers"
	"cloud-native/api/models"
	"cloud-native/api/services/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userSvcMock *mocks.UserService
var userHandler http.Handler

func init() {
	userSvcMock = &mocks.UserService{}
	userHandler = handlers.NewUserHandler(userSvcMock)

}

func TestUserHandler_CreateUser_BadRequestBody(t *testing.T) {
	req := httptest.NewRequest("POST", "/", strings.NewReader(""))
	rec := httptest.NewRecorder()

	userHandler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "invalid request body")

}

func TestCreateUser_UsernameTooShort(t *testing.T) {
	req := httptest.NewRequest("POST", "/", strings.NewReader(`{"username":"us"}`))
	rec := httptest.NewRecorder()

	userHandler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Username must be at least 3 characters.")

}

func TestCreateUser_UsernameGood(t *testing.T) {
	setupMocks()
	req := httptest.NewRequest("POST", "/", strings.NewReader(`{"username":"brian"}`))
	rec := httptest.NewRecorder()

	userSvcMock.On("CreateAccount", mock.Anything).Return("123", nil)

	userHandler.ServeHTTP(rec, req)

	userSvcMock.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "123")

}

func TestCreateUser_CreateAccountError(t *testing.T) {
	setupMocks()
	req := httptest.NewRequest("POST", "/", strings.NewReader(`{"username":"brian"}`))
	rec := httptest.NewRecorder()

	userSvcMock.
		On("CreateAccount", mock.Anything).
		Return("", &models.Error{Code: 400, Message: "username exists"}, nil)

	userHandler.ServeHTTP(rec, req)

	userSvcMock.AssertExpectations(t)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "username exists")

}

func setupMocks() {
	userSvcMock = &mocks.UserService{}
	userHandler = handlers.NewUserHandler(userSvcMock)
}
