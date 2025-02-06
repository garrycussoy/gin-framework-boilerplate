package handlers

import (
	"github.com/gin-gonic/gin"
)

type GeneralHandler struct{}

func NewGeneralHandler() GeneralHandler {
	return GeneralHandler{}
}

// A handler for general health-check
func (generalH GeneralHandler) HealthCheck(ctx *gin.Context) {
	SuccessResponse(ctx, nil)
}
