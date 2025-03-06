package helpers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"gin-framework-boilerplate/pkg/helpers"

	"github.com/stretchr/testify/assert"
)

func TestConvertJSONStringToMapStringInterface(t *testing.T) {
	t.Run("Test 1 | Success", func(t *testing.T) {
		initialStr := "{\"key1\": \"value1\", \"key2\": \"value2\"}"
		convertedMap, err := helpers.ConvertJSONStringToMapStringInterface(initialStr)

		// Assertions
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		}, convertedMap)
	})

	t.Run("Test 2 | Invalid format", func(t *testing.T) {
		initialStr := "{\"key1\"}"
		_, err := helpers.ConvertJSONStringToMapStringInterface(initialStr)

		// Assertions
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "invalid character '}'")
	})
}

func TestConvertStreamToMapStringInterface(t *testing.T) {
	t.Run("Test 1 | Success", func(t *testing.T) {
		initialMap := map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		}
		marshalledJson, _ := json.Marshal(initialMap)
		rdr1 := io.NopCloser(bytes.NewBuffer(marshalledJson))
		convertedMap, err := helpers.ConvertStreamToMapStringInterface(rdr1)

		// Assertions
		assert.Nil(t, err)
		assert.Equal(t, initialMap, convertedMap)
	})

	t.Run("Test 2 | Can't unmarshal", func(t *testing.T) {
		initialStr := "string"
		marshalledJson, _ := json.Marshal(initialStr)
		rdr1 := io.NopCloser(bytes.NewBuffer(marshalledJson))
		_, err := helpers.ConvertStreamToMapStringInterface(rdr1)

		// Assertions
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "cannot unmarshal")
	})
}

// Test ConvertInterfaceToJSONString function
func TestConvertInterfaceToJSONString(t *testing.T) {
	t.Run("Test 1 | Success", func(t *testing.T) {
		initialMap := map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		}
		convertedStr := helpers.ConvertInterfaceToJSONString(initialMap)

		// Assertions
		assert.NotNil(t, convertedStr)
		assert.Equal(t, "{\"key1\":\"value1\",\"key2\":\"value2\"}", convertedStr)
	})
}
