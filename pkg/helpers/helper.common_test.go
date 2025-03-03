package helpers_test

import (
	"testing"

	"gin-framework-boilerplate/pkg/helpers"

	"github.com/stretchr/testify/assert"
)

// Test IsArrayContains function
func TestIsArrayContains(t *testing.T) {
	// Test case 1
	arr := []string{"hello", "world", "golang"}
	str := "golang"
	expected := true
	result := helpers.IsArrayContains(arr, str)

	// Assertions
	assert.Equal(t, expected, result)

	// Test case 2
	arr = []string{"hello", "world", "golang"}
	str = "java"
	expected = false
	result = helpers.IsArrayContains(arr, str)

	// Assertions
	assert.Equal(t, expected, result)
}

// Test GenerateRandomString function
func TestGenerateRandomString(t *testing.T) {
	// Test case 1
	result, err := helpers.GenerateRandomString(48)

	// Assertions
	assert.Len(t, result, 48)
	assert.Nil(t, err)
}

// Test RemoveEmptyField function
func TestRemoveEmptyField(t *testing.T) {
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
}

// Test GenerateUUID function
func TestGenerateUUID(t *testing.T) {
	// Call the function
	uuid, err := helpers.GenerateUUID()

	// Assertions
	assert.Nil(t, err)
	assert.NotEqual(t, "", uuid)
}

// Test ExtractNullableString function
func TestExtractNullableString(t *testing.T) {
	// Test case 1
	var nilStr *string = nil
	extracted := helpers.ExtractNullableString(nilStr)

	// Assertions
	assert.Equal(t, "", extracted)

	// Test case 1
	notNilStr := "This is a string"
	extracted = helpers.ExtractNullableString(&notNilStr)

	// Assertions
	assert.Equal(t, notNilStr, extracted)
}

// Test CreatePointerString function
func TestCreatePointerString(t *testing.T) {
	// Test case 1
	convertedStr := helpers.CreatePointerString("manual string")
	pointerStr := "pointer string"

	// Assertions
	assert.IsType(t, &pointerStr, convertedStr)
}
