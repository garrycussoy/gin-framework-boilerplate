package handlers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	Domains "gin-framework-boilerplate/internal/business/domains"
	Usecases "gin-framework-boilerplate/internal/business/usecases"
	Handlers "gin-framework-boilerplate/internal/http/handlers"
	"gin-framework-boilerplate/internal/mocks"
	ESBPorts "gin-framework-boilerplate/internal/ports/clients/esb"
	Records "gin-framework-boilerplate/internal/ports/repository/records"
	"gin-framework-boilerplate/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userUsecase Domains.UserUsecase
	userHandler Handlers.UserHandler
)

func userTestSetup(t *testing.T) {
	// Define some services
	userRepoMock = mocks.NewUserRepository(t)
	esbClientMock = mocks.NewESBClient(t)
	jwtServiceMock = mocks.NewJWTService(t)
	userUsecase = Usecases.NewUserUsecase(userRepoMock, esbClientMock)
	userHandler = Handlers.NewUserHandler(userUsecase)

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

	userRecordList1 = []Records.User{
		userRecord1,
	}

	esbGeneralResponseDto = ESBPorts.GeneralResponseDTO{
		Message: "Success",
		Code:    "200",
		Data:    nil,
	}

	// Create gin engine
	s = gin.Default()
}

func TestGetUserByEmail(t *testing.T) {
	userTestSetup(t)

	// Define route
	s.GET("/user/:email", userHandler.GetUserByEmail)
	t.Run("Test 1 | Success", func(t *testing.T) {
		// Mocking some functions
		esbClientMock.Mock.On("Sample", mock.Anything).Return(esbGeneralResponseDto, nil).Once()
		userRepoMock.Mock.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(userRecord1, nil).Once()

		// Do some setup
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/user/gin@example.com", nil)
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
		esbClientMock.Mock.On("Sample", mock.Anything).Return(esbGeneralResponseDto, nil).Once()
		userRepoMock.Mock.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(Records.User{}, errors.New("sql: no rows in result set")).Once()

		// Do some setup
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/user/python@example.com", nil)
		r.Header.Set("Content-Type", "application/json")

		// Perform request
		s.ServeHTTP(w, r)

		var body Handlers.BaseResponse
		json.Unmarshal(w.Body.Bytes(), &body)

		// Assertions
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
		assert.Contains(t, w.Result().Header.Get("Content-Type"), "application/json")
		assert.Contains(t, body.Message, "User domain error")
	})

	t.Run("Test 3 | Validation failed (invalid email format)", func(t *testing.T) {
		// Do some setup
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/user/gin", nil)
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

func TestGetUser(t *testing.T) {
	userTestSetup(t)

	// Define route
	s.GET("/users", userHandler.GetUsers)
	t.Run("Test 1 | Success", func(t *testing.T) {
		// Mocking some functions
		userRepoMock.Mock.On("GetUsers", mock.Anything, mock.AnythingOfType("UserFilterDto")).Return(userRecordList1, nil).Once()

		// Do some setup
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/users?branch_id=1&start=2000-01-01&end=2020-01-01", nil)
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

	t.Run("Test 2 | Usecase failed (something's wrong happened)", func(t *testing.T) {
		// Mocking some functions
		userRepoMock.Mock.On("GetUsers", mock.Anything, mock.AnythingOfType("UserFilterDto")).Return([]Records.User{}, errors.New("sql: expected 8 arguments, got 0")).Once()

		// Do some setup
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/users?branch_id=1&start=2000-01-01&end=2020-01-01", nil)
		r.Header.Set("Content-Type", "application/json")

		// Perform request
		s.ServeHTTP(w, r)

		var body Handlers.BaseResponse
		json.Unmarshal(w.Body.Bytes(), &body)

		// Assertions
		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
		assert.Contains(t, w.Result().Header.Get("Content-Type"), "application/json")
		assert.Contains(t, body.Message, "User repository error")
	})
}
