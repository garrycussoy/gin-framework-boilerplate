package middlewares

import (
	"net/http"
	"strings"

	"gin-framework-boilerplate/internal/config"
	"gin-framework-boilerplate/internal/constants"
	"gin-framework-boilerplate/pkg/helpers"
	"gin-framework-boilerplate/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Setup CORS policy
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Define some headers
		c.Writer.Header().Set("Access-Control-Allow-Origin", config.AppConfig.AllowedCORS)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", constants.AllowCredential)
		c.Writer.Header().Set("Access-Control-Allow-Headers", constants.AllowedHeader)
		c.Writer.Header().Set("Access-Control-Allow-Methods", constants.AllowedMethods)
		c.Writer.Header().Set("Access-Control-Max-Age", constants.MaxAge)

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		// Validate method
		if !helpers.IsArrayContains(strings.Split(constants.AllowedMethods, ", "), c.Request.Method) {
			logger.InfoF("method %s is not allowed\n", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryCORS}, c.Request.Method)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden method with CORS policy"})
			return
		}

		// Validate header
		for key, value := range c.Request.Header {
			if !helpers.IsArrayContains(strings.Split(constants.AllowedHeader, ", "), key) {
				logger.InfoF("%s: %s\n", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryCORS}, key, value)
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden with CORS policy"})
				return
			}
		}

		// Validate origin
		if config.AppConfig.AllowedCORS != "*" {
			if !helpers.IsArrayContains(strings.Split(config.AppConfig.AllowedCORS, ", "), c.Request.Host) {
				logger.InfoF("host '%s' is not part of '%v'\n", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryCORS}, c.Request.Host, config.AppConfig.AllowedCORS)
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden host with CORS policy"})
				return
			}
		}

		c.Next()
	}
}
