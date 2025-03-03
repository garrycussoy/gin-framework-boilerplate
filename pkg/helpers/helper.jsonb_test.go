package helpers_test

import (
	"testing"

	"gin-framework-boilerplate/pkg/helpers"

	"github.com/stretchr/testify/assert"
)

// Test JSONB functionality
func TestJSONBMethod(t *testing.T) {
	// Test Case 1 : Success
	// Define a JSONB object
	data := helpers.JSONB{
		"email":    "jsonb@example.com",
		"username": "jsonb",
	}

	// Let's try marshalling the data
	marshalledData, err := data.Value()
	assert.Nil(t, err)

	// Define another JSONB object
	var otherData helpers.JSONB

	// Scan the marshalled data
	err = otherData.Scan(marshalledData)
	assert.Nil(t, err)

	// Check the type
	email, ok := otherData["email"].(string)
	assert.True(t, ok)
	if ok {
		assert.Equal(t, "jsonb@example.com", email)
	}

	// Test Case 2 : The argument for Scan() is nil
	err = otherData.Scan(nil)
	assert.Nil(t, err)

	// Test Case 3 : The argument for Scan() isn't a byte[] format
	err = otherData.Scan("not a byte[]")
	assert.NotNil(t, err)
	assert.Equal(t, "type assertion to []byte failed", err.Error())
}
