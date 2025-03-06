package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	Domains "gin-framework-boilerplate/internal/business/domains"
	Usecases "gin-framework-boilerplate/internal/business/usecases"
	Requests "gin-framework-boilerplate/internal/http/datatransfers/requests"
	Handlers "gin-framework-boilerplate/internal/http/handlers"
	"gin-framework-boilerplate/internal/mocks"
	Records "gin-framework-boilerplate/internal/ports/repository/records"
	"gin-framework-boilerplate/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	authUsecase Domains.AuthUsecase
	authHandler Handlers.AuthHandler
)

func authTestSetup(t *testing.T) {
	// Define some services
	jwtServiceMock = mocks.NewJWTService(t)
	userRepoMock = mocks.NewUserRepository(t)
	authUsecase = Usecases.NewAuthUsecase(jwtServiceMock, userRepoMock)
	authHandler = Handlers.NewAuthHandler(authUsecase)

	// Define some variables
	userRecord1 = Records.User{
		Id:          "c5332d52-933c-4cd1-9a0c-ed88b232cc66",
		FullName:    "User 1",
		Email:       "gin@example.com",
		PhoneNumber: "089508960001",
		Password:    "$2a$10$7tMnIqKUaVLlQygWpJssduxQoQpK7ZZYI1e/RtnTAY27An3aDG7Bq",
		BranchId:    helpers.CreatePointerString("1"),
		CreatedAt:   time.Now(),
	}

	// Create gin engine
	s = gin.Default()
}

func TestUserLogin(t *testing.T) {
	authTestSetup(t)

	// Define route
	s.POST("/login", authHandler.UserLogin)
	t.Run("Test 1 | Success", func(t *testing.T) {
		// Mocking some functions
		userRepoMock.Mock.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(userRecord1, nil).Once()
		jwtServiceMock.Mock.On("GenerateToken", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("token", nil).Once()

		// Prepare payload
		req := Requests.UserLoginRequest{
			Email:    "gin@example.com",
			Password: "Password365!",
		}
		reqBody, _ := json.Marshal(req)

		// Do some setup
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(reqBody))
		r.Header.Set("Content-Type", "application/json")

		// Perform request
		s.ServeHTTP(w, r)

		var body Handlers.BaseResponse
		json.Unmarshal(w.Body.Bytes(), &body)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Contains(t, w.Result().Header.Get("Content-Type"), "application/json")
		assert.Contains(t, body.Message, "Success")
	})

	t.Run("Test 2 | Usecase failed (email isn't registered)", func(t *testing.T) {
		// Mocking some functions
		userRepoMock.Mock.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(userRecord1, errors.New("sql: no rows in result set")).Once()

		// Prepare payload
		req := Requests.UserLoginRequest{
			Email:    "echo@example.com",
			Password: "Password365!",
		}
		reqBody, _ := json.Marshal(req)

		// Do some setup
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(reqBody))
		r.Header.Set("Content-Type", "application/json")

		// Perform request
		s.ServeHTTP(w, r)

		var body Handlers.BaseResponse
		json.Unmarshal(w.Body.Bytes(), &body)

		// Assertions
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
		assert.Contains(t, w.Result().Header.Get("Content-Type"), "application/json")
		assert.Contains(t, body.Message, "Auth domain error")
	})

	t.Run("Test 3 | Validation failed (missing required field)", func(t *testing.T) {
		// Prepare payload
		req := Requests.UserLoginRequest{
			Email: "gin@example.com",
		}
		reqBody, _ := json.Marshal(req)

		// Do some setup
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(reqBody))
		r.Header.Set("Content-Type", "application/json")

		// Perform request
		s.ServeHTTP(w, r)

		var body Handlers.BaseResponse
		json.Unmarshal(w.Body.Bytes(), &body)

		// Assertions
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
		assert.Contains(t, w.Result().Header.Get("Content-Type"), "application/json")
		assert.Contains(t, body.Message, "Payload validation failed")
	})

	t.Run("Test 4 | Binding error", func(t *testing.T) {
		// Prepare payload
		req := "invalid-payload"
		reqBody, _ := json.Marshal(req)

		// Do some setup
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(reqBody))
		r.Header.Set("Content-Type", "application/json")

		// Perform request
		s.ServeHTTP(w, r)

		var body Handlers.BaseResponse
		json.Unmarshal(w.Body.Bytes(), &body)

		// Assertions
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
		assert.Contains(t, w.Result().Header.Get("Content-Type"), "application/json")
		assert.Contains(t, body.Message, "Payload validation failed")
	})
}
