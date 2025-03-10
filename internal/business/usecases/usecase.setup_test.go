package usecases_test

import (
	"gin-framework-boilerplate/internal/mocks"

	ESBPorts "gin-framework-boilerplate/internal/ports/clients/esb"
	Records "gin-framework-boilerplate/internal/ports/repository/records"

	"github.com/gin-gonic/gin"
)

// Define some variables which will be used accross usecase tests
var (
	// General
	s *gin.Engine

	// Mock repositories
	userRepoMock *mocks.UserRepository

	// Mock service
	jwtServiceMock *mocks.JWTService
	esbClientMock  *mocks.ESBClient

	// Records
	userRecord1     Records.User
	userRecordList1 []Records.User

	// Client DTO
	esbGeneralResponseDto ESBPorts.GeneralResponseDTO
)
