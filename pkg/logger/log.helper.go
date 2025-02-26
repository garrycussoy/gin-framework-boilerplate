package logger

import (
	"fmt"
	"gin-framework-boilerplate/pkg/helpers"
)

// Define which fields need to be masked
var sensitiveFields = []string{
	"password",
	"new_password",
	"confirm_password",
	"token",
	"access_token",
}

// A function to masked sensitive values
func MaskingValues(log map[string]interface{}) map[string]interface{} {
	for k, v := range log {
		// Check for nested case
		converted, isNested := v.(map[string]interface{})
		if isNested && converted != nil {
			// If the code goes here, then it is nested case
			MaskingValues(converted)
		}

		// Mask sensitive fields
		if helpers.IsArrayContains(sensitiveFields, k) {
			log[k] = "*****"
		}
	}
	return log
}

// A function to formatting form-data
func FormattingFormData(log map[string]interface{}) map[string]interface{} {
	for k, v := range log {
		// Turn the value into string
		valueStr := fmt.Sprintf("%s", v)
		log[k] = valueStr[1 : len(valueStr)-1]
	}
	return log
}
