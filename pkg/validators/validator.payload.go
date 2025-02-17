// This package provide validator function to validate the payload of incoming request
package validators

import (
	"errors"
	"fmt"
	"strings"

	"gin-framework-boilerplate/pkg/helpers"

	"github.com/go-playground/validator/v10"
)

// Following map will show basic error message for each criteria
var mapHelper = map[string]string{
	"required":  "is a required field",
	"email":     "is not a valid email address",
	"lowercase": "must contain at least one lowercase letter",
	"uppercase": "must contain at least one uppercase letter",
	"numeric":   "must contain at least one digit",
}

// Define some tags which need parameter
var needParam = []string{"min", "max", "containsany", "required_if"}

// Function to validate the payload
func ValidatePayloads(payload interface{}) (err error) {
	// Define validator instance and some variables
	validate := validator.New()
	var field, param, tag, message string

	// Do struct validation
	err = validate.Struct(payload)
	if err != nil {
		// We'll loop through each error detected (but only show one of the error at a time)
		for _, e := range err.(validator.ValidationErrors) {
			// Get some basic error info
			field = e.Field()
			tag = e.Tag()
			param = e.Param()

			// Check whether the error is one in the tag which need param or not
			if helpers.IsArrayContains(needParam, tag) {
				message = errWithParam(field, tag, param)
				continue
			}

			// Formatting error message
			message = fmt.Sprintf("%s: %s", field, mapHelper[tag])
		}

		return errors.New(message)
	}

	return nil
}

// Function to format error message
func errWithParam(field, tag, param string) string {
	// Format message based on tag
	var message string
	switch tag {
	case "min":
		message = fmt.Sprintf("must be at least %s characters long", param)
	case "max":
		message = fmt.Sprintf("must be less than %s characters", param)
	case "containsany":
		message = fmt.Sprintf("must contain at least one symbol of '%s'", param)
	case "required_if":
		message = fmt.Sprintf("is required since '%s' equals '%s'", strings.Split(param, " ")[0], strings.Split(param, " ")[1])
	}

	return fmt.Sprintf("%s: %s", field, message)
}
