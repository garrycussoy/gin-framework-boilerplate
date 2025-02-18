package middlewares

import (
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"

	"gin-framework-boilerplate/internal/config"
	"gin-framework-boilerplate/internal/http/handlers"
	custom_errors "gin-framework-boilerplate/pkg/errors"
)

// Display timeout error message
func timeoutErrorResponse(c *gin.Context) {
	handlers.ErrorResponse(c, custom_errors.TimeoutLimitExceeded())
}

// Middleware which will handle endpoint's timeout
func TimeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(time.Duration(config.AppConfig.HandlerTimeout)*time.Millisecond),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(timeoutErrorResponse),
	)
}
