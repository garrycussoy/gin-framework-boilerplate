// This package is used to create new object called JSONB
// This object will be used to communicate between system and database for JSONB object type
package helpers

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Define JSONB type
type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONB) Scan(value interface{}) error {
	// Do the scanning only if the value not nil
	if value != nil {
		b, ok := value.([]byte)
		if !ok {
			return errors.New("type assertion to []byte failed")
		}

		return json.Unmarshal(b, &j)
	}

	return nil
}
