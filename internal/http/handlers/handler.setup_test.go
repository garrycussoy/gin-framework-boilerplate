package handlers_test

import (
	"gin-framework-boilerplate/internal/mocks"

	Records "gin-framework-boilerplate/internal/ports/repository/records"

	"github.com/gin-gonic/gin"
)

// Define some variables which will be used accross usecase tests
var (
	// Mock service
	userRepoMock   *mocks.UserRepository
	jwtServiceMock *mocks.JWTService
	s              *gin.Engine

	// Records
	userRecord1 Records.User
)
