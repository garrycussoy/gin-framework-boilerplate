package logger

import (
	"bytes"
	"fmt"
	"io"

	"gin-framework-boilerplate/internal/constants"
	"gin-framework-boilerplate/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Define color constants
const (
	Red    = "41"
	Yellow = "43"
	Green  = "42"
)

// Define which methods that will trigger PayloadRequestLogger flow
var allowedMethods = []string{
	"POST",
	"PUT",
	"DELETE",
}

// Function to log some basic data of incoming HTTP request
func HTTPLogger(param gin.LogFormatterParams) string {
	// Adjust color in the ouptut log
	var color string
	switch {
	case param.StatusCode >= 500:
		color = Red
	case param.StatusCode >= 400:
		color = Yellow
	default:
		color = Green
	}

	// Formatting the log
	var logFormat = fmt.Sprintf("%s \033[%sm %d \033[0m %s %s %d %s %s %s CorrelationID: %s\n",
		constants.HTTPLogging,
		color,
		param.StatusCode,
		param.Method,
		param.Path,
		param.Latency,
		param.ClientIP,
		param.ErrorMessage,
		param.Request.UserAgent(),
		param.Keys["CorrelationID"],
	)

	// Output the log
	Info(logFormat, logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})

	return ""
}

// A function to log incoming HTTP request payload
func PayloadRequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Do following process only for allowed methods
		if helpers.IsArrayContains(allowedMethods, c.Request.Method) {
			buf, _ := io.ReadAll(c.Request.Body)
			rdr1 := io.NopCloser(bytes.NewBuffer(buf))
			rdr2 := io.NopCloser(bytes.NewBuffer(buf)) // We have to create a new Buffer, because rdr1 will be read

			// Formatting the log
			bodyRequest, err := ReadBody(rdr1)
			if err != nil {
				// For debugging puprose only
				logFormat := fmt.Sprintf("%s Error while marshalling body request for %s %s. Detail: %s. CorrelationId: %s\n",
					constants.HTTPLogging,
					c.Request.Method,
					c.Request.URL.Path,
					err.Error(),
					c.GetString("CorrelationID"),
				)
				Debug(logFormat, logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
			} else {
				logFormat := fmt.Sprintf("%s %s %s BodyRequest: %v. CorrelationId: %s\n",
					constants.HTTPLogging,
					c.Request.Method,
					c.Request.URL.Path,
					bodyRequest,
					c.GetString("CorrelationID"),
				)
				Info(logFormat, logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
			}

			c.Request.Body = rdr2
		}
		c.Next()
	}
}
