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

// Define which methods that will trigger RequestPayloadLogger flow
var allowedMethods = []string{
	"POST",
	"PUT",
	"DELETE",
}

// Define custom body log writer
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// We need to redefine Write method
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
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
	logFormat := map[string]string{
		"CorrelationID": fmt.Sprintf("%s", param.Keys["CorrelationId"]),
		"Latency":       param.Latency.String(),
		"ClientIP":      param.ClientIP,
		"UserAgent":     param.Request.UserAgent(),
	}

	// Add errorMessage only if the value isn't empty
	if param.ErrorMessage != "" {
		logFormat["ErrorMessage"] = param.ErrorMessage
	}

	var message = fmt.Sprintf("%s \033[%sm %d \033[0m %s %s %v\n",
		constants.HTTPLogging,
		color,
		param.StatusCode,
		param.Method,
		param.Path,
		helpers.ConvertInterfaceToJSONString(logFormat),
	)

	// Output the log
	Info(message, logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})

	return ""
}

// A function to log incoming HTTP request payload
func RequestPayloadLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Do following process only for allowed methods
		if helpers.IsArrayContains(allowedMethods, c.Request.Method) {
			buf, _ := io.ReadAll(c.Request.Body)
			rdr1 := io.NopCloser(bytes.NewBuffer(buf))
			rdr2 := io.NopCloser(bytes.NewBuffer(buf)) // We have to create a new Buffer, because rdr1 will be read

			// Read the body, and turn it into a map
			bodyRequest, err := helpers.ConvertStreamToMapStringInterface(rdr1)
			if err != nil {
				// For debugging puprose only
				// Define log content
				logContent := map[string]interface{}{
					"Message":       fmt.Sprintf("Error while marshalling body request for %s %s", c.Request.Method, c.Request.URL.Path),
					"Detail":        err.Error(),
					"CorrelationId": c.GetString("CorrelationID"),
				}

				// Formatting the log
				logFormat := fmt.Sprintf("%s %v\n",
					constants.HTTPLogging,
					helpers.ConvertInterfaceToJSONString(logContent),
				)
				Debug(logFormat, logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
			} else {
				// Define log content
				logContent := map[string]interface{}{
					"CorrelationId": c.GetString("CorrelationId"),
					"BodyRequest":   bodyRequest,
				}

				// Masking some sensitive values
				logContent = MaskingValues(logContent)

				// Formatting the log
				logFormat := fmt.Sprintf("%s %s %s %v\n",
					constants.HTTPLogging,
					c.Request.Method,
					c.Request.URL.Path,
					helpers.ConvertInterfaceToJSONString(logContent),
				)
				Info(logFormat, logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
			}

			c.Request.Body = rdr2
		}
		c.Next()
	}
}

// A function to log the response payload of incoming HTTP request
func ResponsePayloadLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get correlationId
		logContent := map[string]interface{}{
			"CorrelationId": c.GetString("CorrelationId"),
		}

		// Create new instance of body log writer
		writer := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = writer
		c.Next()

		// Get response body
		responseBody, _ := helpers.ConvertJSONStringToMapStringInterface(writer.body.String())
		logContent["ResponseBody"] = responseBody

		// Masking some sensitive values
		logContent = MaskingValues(logContent)

		// Formatting the log
		logFormat := fmt.Sprintf("%s %s %s %v\n",
			constants.HTTPLogging,
			c.Request.Method,
			c.Request.URL.Path,
			helpers.ConvertInterfaceToJSONString(logContent),
		)

		Info(logFormat, logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})
	}
}
