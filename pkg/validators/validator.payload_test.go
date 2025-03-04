package validators_test

import (
	"testing"

	Validator "gin-framework-boilerplate/pkg/validators"

	"github.com/stretchr/testify/assert"
)

// Define some payloads needed for the test
type DummyPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// Test ValidatePayloads function
func TestValidatePayloads(t *testing.T) {
	// Test case 1: detect missing value of a required field
	payload := DummyPayload{
		Email:    "gin@example.com",
		Password: "",
	}
	err := Validator.ValidatePayloads(payload)

	// Assertions
	assert.NotNil(t, err)
	assert.Equal(t, "Password: is a required field", err.Error())

	// Test case 2: detect invalid format
	payload = DummyPayload{
		Email:    "this.is-not_an.email",
		Password: "password876@",
	}
	err = Validator.ValidatePayloads(payload)

	// Assertions
	assert.NotNil(t, err)
	assert.Equal(t, "Email: is not a valid email address", err.Error())

	// Test case 3: error with params
	payload = DummyPayload{
		Email:    "gin@example.com",
		Password: "short",
	}
	err = Validator.ValidatePayloads(payload)

	// Assertions
	assert.NotNil(t, err)
	assert.Equal(t, "Password: must be at least 6 characters long", err.Error())

	// Test case 4: validation passed
	payload = DummyPayload{
		Email:    "gin@example.com",
		Password: "Password999!",
	}
	err = Validator.ValidatePayloads(payload)

	// Assertions
	assert.Nil(t, err)
}
