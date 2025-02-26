package logger

import (
	"fmt"
	"gin-framework-boilerplate/internal/constants"
	"gin-framework-boilerplate/pkg/helpers"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// A function to log external HTTP request
func ExternalHTTPRequestLogger(c *resty.Client, req *resty.Request) error {
	// Differentiate some logging data based on method
	var request interface{}
	var requestFormat string
	if req.Method != "GET" {
		// Look for body request or form data
		if req.Body != nil {
			request = req.Body
			requestFormat = "BodyRequest"
		} else if req.FormData != nil {
			request = req.FormData
			requestFormat = "FormData"
		}
	} else {
		requestFormat = "QueryParam"

		// Convert query param into map[string]interface{}
		mappedQueryParam := make(map[string]interface{})
		for k, v := range req.QueryParam {
			mappedQueryParam[k] = v[0]
		}
		request = mappedQueryParam
	}

	// Convert body request into map format
	bodyRequest, _ := helpers.ConvertJSONStringToMapStringInterface(helpers.ConvertInterfaceToJSONString(request))

	// Differentiate some logging format based on method
	if req.Method != "GET" {
		if req.Body == nil {
			// Formatting form data
			bodyRequest = FormattingFormData(bodyRequest)
		}
	}

	// Masking sensitive fields
	bodyRequest = MaskingValues(bodyRequest)

	// Formatting the log
	logFormat := map[string]interface{}{
		"CorrelationID": "TEST-001",
		requestFormat:   bodyRequest,
	}

	// Logging
	message := fmt.Sprintf("%s %s %s %s\n",
		constants.ExternalHTTPLogging,
		req.Method,
		req.URL,
		helpers.ConvertInterfaceToJSONString(logFormat),
	)

	// Output the message
	Info(message, logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})

	return nil
}

// A function to log external HTTP response
func ExternalHTTPResponseLogger(c *resty.Client, resp *resty.Response) error {
	// Convert response data into map format
	respData, _ := helpers.ConvertJSONStringToMapStringInterface(helpers.ConvertInterfaceToJSONString(resp.Request.Result))

	// Masking some sensitive fields
	respData = MaskingValues(respData)

	// Formatting the log
	logFormat := map[string]interface{}{
		"CorrelationID": "TEST-001",
		"ResponseData":  respData,
	}

	message := fmt.Sprintf("%s %s %s %s %s\n",
		constants.ExternalHTTPLogging,
		resp.Status(),
		resp.Request.Method,
		resp.Request.URL,
		helpers.ConvertInterfaceToJSONString(logFormat),
	)

	// Output the message
	Info(message, logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryHTTP})

	return nil
}
