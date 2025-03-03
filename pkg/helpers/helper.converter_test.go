package helpers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"gin-framework-boilerplate/pkg/helpers"

	"github.com/stretchr/testify/assert"
)

// Test ConvertJSONStringToMapStringInterface function
func TestConvertJSONStringToMapStringInterface(t *testing.T) {
	// Test case 1 (Success)
	initialStr := "{\"key1\": \"value1\", \"key2\": \"value2\"}"
	convertedMap, err := helpers.ConvertJSONStringToMapStringInterface(initialStr)

	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}, convertedMap)

	// Test case 2 (Error)
	initialStr = "{\"key1\"}"
	_, err = helpers.ConvertJSONStringToMapStringInterface(initialStr)

	// Assertions
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "invalid character '}'")
}

// Test ConvertStreamToMapStringInterface function
func TestConvertStreamToMapStringInterface(t *testing.T) {
	// Test case 1 (Success)
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

	// Test case 1 (Error)
	initialStr := "string"
	marshalledJson, _ = json.Marshal(initialStr)
	rdr1 = io.NopCloser(bytes.NewBuffer(marshalledJson))
	_, err = helpers.ConvertStreamToMapStringInterface(rdr1)

	// Assertions
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "cannot unmarshal")
}

// Test ConvertInterfaceToJSONString function
func TestConvertInterfaceToJSONString(t *testing.T) {
	// Test case 1 (Success)
	initialMap := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}
	convertedStr := helpers.ConvertInterfaceToJSONString(initialMap)

	// Assertions
	assert.NotNil(t, convertedStr)
	assert.Equal(t, "{\"key1\":\"value1\",\"key2\":\"value2\"}", convertedStr)
}
