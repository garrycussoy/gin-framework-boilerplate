package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseSuccessResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// type BasePartialSuccessResponse struct {
// 	Code    string      `json:"code"`
// 	Message string      `json:"message,omitempty"`
// 	Data    interface{} `json:"data,omitempty"`
// 	Detail  interface{} `json:"detail,omitempty"`
// }

// type BaseErrorResponse struct {
// 	Code    string      `json:"code"`
// 	Message string      `json:"message,omitempty"`
// 	Detail  interface{} `json:"detail,omitempty"`
// }

// Default success response
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, BaseSuccessResponse{
		Code:    "0000",
		Message: "Success",
		Data:    data,
	})
}

// func PartialSuccessResponse(c *gin.Context, rcCode string, message string, data interface{}, err interface{}) {
// 	c.JSON(http.StatusOK, BasePartialSuccessResponse{
// 		Code:    rcCode,
// 		Message: message,
// 		Data:    data,
// 		Detail:  err,
// 	})
// }

// func ErrorResponse(c *gin.Context, err Errors.AkasiaError) {
// 	c.JSON(err.Error().Status, BaseErrorResponse{
// 		Code:    err.Error().Code,
// 		Message: err.Error().Message,
// 		Detail:  err.Error().Detail,
// 	})
// }

// func AbortResponse(c *gin.Context, err Errors.AkasiaError) {
// 	c.AbortWithStatusJSON(err.Error().Status, BaseErrorResponse{
// 		Code:    err.Error().Code,
// 		Message: err.Error().Message,
// 		Detail:  err.Error().Detail,
// 	})
// }
