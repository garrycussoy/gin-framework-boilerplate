package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	Errors "gin-framework-boilerplate/pkg/errors"
)

// Response template which will be used for both success and error response
type BaseResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// Function to output success response
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, BaseResponse{
		Code:    "0000",
		Message: "Success",
		Data:    data,
	})
}

// Function to output error response
func ErrorResponse(c *gin.Context, err Errors.CustomError) {
	c.JSON(err.Error().Status, BaseResponse{
		Code:    err.Error().Code,
		Message: err.Error().Message,
		Error:   err.Error().Detail,
	})
}

// Function to output non-blocking error response
func NonBlockingErrorResponse(c *gin.Context, rcCode string, message string, data interface{}, err interface{}) {
	c.JSON(http.StatusOK, BaseResponse{
		Code:    rcCode,
		Message: message,
		Data:    data,
		Error:   err,
	})
}

// Function to abort the process and give an error response
func AbortResponse(c *gin.Context, err Errors.CustomError) {
	c.AbortWithStatusJSON(err.Error().Status, BaseResponse{
		Code:    err.Error().Code,
		Message: err.Error().Message,
		Error:   err.Error().Detail,
	})
}
