package helpers_test

import (
	"testing"

	"gin-framework-boilerplate/pkg/helpers"

	"github.com/stretchr/testify/assert"
)

func TestJSONBMethod(t *testing.T) {
	// Define some objects needed
	data := helpers.JSONB{
		"email":    "jsonb@example.com",
		"username": "jsonb",
	}
	var otherData helpers.JSONB
	var err error

	t.Run("Test 1 | Success marshalling and scanning data", func(t *testing.T) {
		// Let's try marshalling the data
		marshalledData, err := data.Value()
		assert.Nil(t, err)

		// Scan the marshalled data
		err = otherData.Scan(marshalledData)
		assert.Nil(t, err)

		// Check the type
		email, ok := otherData["email"].(string)
		assert.True(t, ok)
		if ok {
			assert.Equal(t, "jsonb@example.com", email)
		}
	})

	t.Run("Test 2 | The argument for Scan() is nil", func(t *testing.T) {
		err = otherData.Scan(nil)
		assert.Nil(t, err)
	})

	t.Run("Test 3 | The argument for Scan() isn't a byte[] format", func(t *testing.T) {
		err = otherData.Scan("not a byte[]")
		assert.NotNil(t, err)
		assert.Equal(t, "type assertion to []byte failed", err.Error())
	})
}
