package usecases_test

import (
	"context"
	"errors"
	"testing"
	"time"

	Domains "gin-framework-boilerplate/internal/business/domains"
	Usecases "gin-framework-boilerplate/internal/business/usecases"
	"gin-framework-boilerplate/internal/mocks"
	Records "gin-framework-boilerplate/internal/ports/repository/records"
	"gin-framework-boilerplate/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	authUsecase Domains.AuthUsecase
)

func authTestSetup(t *testing.T) {
	// Define some services
	jwtServiceMock = mocks.NewJWTService(t)
	userRepoMock = mocks.NewUserRepository(t)
	authUsecase = Usecases.NewAuthUsecase(jwtServiceMock, userRepoMock)

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

	t.Run("Test 1 | Success", func(t *testing.T) {
		// Mocking some functions
		userRepoMock.Mock.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(userRecord1, nil).Once()
		jwtServiceMock.Mock.On("GenerateToken", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("token", nil).Once()

		// Test the function
		result, err := authUsecase.UserLogin(context.Background(), &Domains.UserLoginDomain{
			Email:    "gin@example.com",
			Password: "Password365!",
		})

		// Assertions
		assert.NotNil(t, result)
		assert.Equal(t, "token", result.Token)
		assert.Nil(t, err)
	})

	t.Run("Test 2 | Failed generating token", func(t *testing.T) {
		// Mocking some functions
		userRepoMock.Mock.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(userRecord1, nil).Once()
		jwtServiceMock.Mock.On("GenerateToken", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", errors.New("error while generating token")).Once()

		// Test the function
		_, err := authUsecase.UserLogin(context.Background(), &Domains.UserLoginDomain{
			Email:    "gin@example.com",
			Password: "Password365!",
		})

		// Assertions
		assert.NotNil(t, err)
		assert.Equal(t, 500, err.Error().Status)
		assert.Equal(t, "Auth domain error", err.Error().Message)
	})

	t.Run("Test 3 | Invalid password", func(t *testing.T) {
		// Mocking some functions
		userRepoMock.Mock.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(userRecord1, nil).Once()

		// Test the function
		_, err := authUsecase.UserLogin(context.Background(), &Domains.UserLoginDomain{
			Email:    "gin@example.com",
			Password: "Password111#",
		})

		// Assertions
		assert.NotNil(t, err)
		assert.Equal(t, 400, err.Error().Status)
		assert.Equal(t, "Auth domain error", err.Error().Message)
	})

	t.Run("Test 4 | User not found", func(t *testing.T) {
		// Mocking some functions
		userRepoMock.Mock.On("GetUserByEmail", mock.Anything, mock.AnythingOfType("string")).Return(Records.User{}, errors.New("sql: no rows in result set")).Once()

		// Test the function
		_, err := authUsecase.UserLogin(context.Background(), &Domains.UserLoginDomain{
			Email:    "python@example.com",
			Password: "Password365!",
		})

		// Assertions
		assert.NotNil(t, err)
		assert.Equal(t, 400, err.Error().Status)
		assert.Equal(t, "Auth domain error", err.Error().Message)
	})
}
