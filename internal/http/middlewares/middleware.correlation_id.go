package middlewares

import (
	"github.com/gin-gonic/gin"

	helpers "gin-framework-boilerplate/pkg/helpers"
)

// Setup correlation-ID generator middleware
func CorrelationIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate unique correlation-ID
		corrId, _ := helpers.GenerateUUID()

		// Setup CorrelationId
		c.Set("CorrelationId", corrId)
	}
}
