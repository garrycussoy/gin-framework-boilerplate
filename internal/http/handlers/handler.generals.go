package handlers

import (
	"github.com/gin-gonic/gin"
)

type GeneralHandler struct{}

func NewGeneralHandler() GeneralHandler {
	return GeneralHandler{}
}

// @Summary Basic check of system's health
// @Description Will return 200 if everything works properly.
// @Tags General
// @Produce json
// @Success 200 {object} nil
// @Router /health-check [get]
func (generalH GeneralHandler) HealthCheck(ctx *gin.Context) {
	SuccessResponse(ctx, nil)
}
