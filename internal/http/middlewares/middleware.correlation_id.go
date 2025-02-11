package middlewares

import (
	"github.com/gin-gonic/gin"

	uuid "github.com/nu7hatch/gouuid"
)

// Setup correlation-ID generator middleware
func CorrelationIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate unique correlation-ID
		corrId, _ := uuid.NewV4()

		// Setup CorrelationID
		c.Set("CorrelationID", corrId.String())
	}
}
