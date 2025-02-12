package helpers

import (
	"bytes"
	"encoding/json"
	"io"
)

// Function to read stream and turn it into array of bytes
func ConvertStreamToMapStringInterface(reader io.Reader) (map[string]interface{}, error) {
	// Read the stream
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	// Convert it to string
	s := buf.String()

	// Unmarshal the JSON data
	var result map[string]interface{}
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return map[string]interface{}{}, err
	}

	return result, nil
}

// A function to convert any interface into plain JSON string
func ConvertInterfaceToJSONString(data interface{}) string {
	byteData, err := json.Marshal(data)
	if err != nil {
		// In case something's wrong, we will just print an empty string
		return ""
	}

	return string(byteData)
}
