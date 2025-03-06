package logger_test

import (
	"testing"

	"gin-framework-boilerplate/pkg/logger"

	"github.com/stretchr/testify/assert"
)

func TestFormattingFormData(t *testing.T) {
	t.Run("Test 1 | Success", func(t *testing.T) {
		// Define sample data
		sampleData := map[string]interface{}{
			"key1": "\"value1\"",
			"key2": "\"value2\"",
		}

		// Call the function
		formattedData := logger.FormattingFormData(sampleData)

		// Assertion
		assert.Equal(t, map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		}, formattedData)
	})
}

func TestMaskingValues(t *testing.T) {
	t.Run("Test 1 | Success", func(t *testing.T) {
		// Define sample data
		sampleData := map[string]interface{}{
			"username": "\"john_doe\"",
			"password": "\"Pass987!\"",
		}

		// Call the function
		maskedData := logger.MaskingValues(sampleData)

		// Assertion
		assert.Equal(t, map[string]interface{}{
			"username": "\"john_doe\"",
			"password": "*****",
		}, maskedData)
	})
}
