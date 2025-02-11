package logger

import (
	"bytes"
	"encoding/json"
	"io"

	"gin-framework-boilerplate/internal/constants"
	"gin-framework-boilerplate/pkg/helpers"

	"github.com/sirupsen/logrus"
)

// Define which fields need to be masked
var sensitiveFields = []string{
	"password",
	"new_password",
	"confirm_password",
	"token",
	"access_token",
}

// Read the body message and turn it into json string
func ReadBody(reader io.Reader) (string, error) {
	// Read the body
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	s := buf.String()

	// Unmarshal the JSON data
	var result map[string]interface{}
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return "", err
	}

	// Mask sensitive values
	result = MaskingValues(result)

	// Print as a string
	return ConvertToJSONString(result), nil
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

// A function to convert any interface into plain JSON string
func ConvertToJSONString(data interface{}) string {
	byteData, err := json.Marshal(data)
	if err != nil {
		Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})

		// In case something's wrong, we will just print an empty string
		return ""
	}

	return string(byteData)
}

// A function to formatting form-data
// func FormattingFormData(log map[string]interface{}) map[string]interface{} {
// 	for k, v := range log {
// 		// Turn the value into string
// 		valueStr := fmt.Sprintf("%s", v)
// 		log[k] = valueStr[1 : len(valueStr)-1]
// 	}
// 	return log
// }

// // A function to convert any interface into map[string]interface{} format
// func ConvertToMapStringInterface(data interface{}) map[string]interface{} {
// 	var converted map[string]interface{}
// 	byteData, err := json.Marshal(data)
// 	if err != nil {
// 		Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})

// 		// In case something's wrong, we will just print an empty map[string]interface{}
// 		return make(map[string]interface{})
// 	}

// 	err = json.Unmarshal(byteData, &converted)
// 	if err != nil {
// 		Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})

// 		// In case something's wrong, we will just print an empty map[string]interface{}
// 		return make(map[string]interface{})
// 	}

// 	return converted
// }
