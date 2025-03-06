package helpers_test

import (
	"testing"

	"gin-framework-boilerplate/pkg/helpers"

	"github.com/stretchr/testify/assert"
)

// Test IsArrayContains function
func TestIsArrayContains(t *testing.T) {
	arr := []string{"hello", "world", "golang"}
	tests := []struct {
		name     string
		str      string
		expected bool
	}{
		{"Test 1 | Contains given element", "golang", true},
		{"Test 2 | Doesn't contain given element", "python", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := helpers.IsArrayContains(arr, test.str)

			// Assertions
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestGenerateRandomString(t *testing.T) {
	t.Run("Test 1 | Success generating random string", func(t *testing.T) {
		result, err := helpers.GenerateRandomString(48)

		// Assertions
		assert.Len(t, result, 48)
		assert.Nil(t, err)
	})
}

func TestRemoveEmptyField(t *testing.T) {
	t.Run("Test 1 | Success removing empty field", func(t *testing.T) {
		// Define sample data
		sampleData := map[string]string{
			"first_name": "John",
			"last_name":  "",
			"gender":     "1",
			"dob":        "2000-09-08",
		}

		// Call the function
		formatted := helpers.RemoveEmptyField(sampleData)

		// Assertions
		assert.Equal(t, map[string]string{
			"first_name": "John",
			"gender":     "1",
			"dob":        "2000-09-08",
		}, formatted)
	})
}

func TestGenerateUUID(t *testing.T) {
	t.Run("Test 1 | Success generating UUID", func(t *testing.T) {
		// Call the function
		uuid, err := helpers.GenerateUUID()

		// Assertions
		assert.Nil(t, err)
		assert.NotEqual(t, "", uuid)
	})
}

func TestExtractNullableString(t *testing.T) {
	t.Run("Test 1 | Nil string", func(t *testing.T) {
		var nilStr *string = nil
		extracted := helpers.ExtractNullableString(nilStr)

		// Assertion
		assert.Equal(t, "", extracted)
	})

	t.Run("Test 2 | Not a nil string", func(t *testing.T) {
		notNilStr := "This is a string"
		extracted := helpers.ExtractNullableString(&notNilStr)

		// Assertion
		assert.Equal(t, notNilStr, extracted)
	})
}

func TestCreatePointerString(t *testing.T) {
	t.Run("Test 1 | Success creating a pointer string", func(t *testing.T) {
		convertedStr := helpers.CreatePointerString("manual string")
		pointerStr := "pointer string"

		// Assertion
		assert.IsType(t, &pointerStr, convertedStr)
	})
}
