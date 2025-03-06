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

func TestValidatePayloads(t *testing.T) {
	t.Run("Test 1 | Missing required field", func(t *testing.T) {
		payload := DummyPayload{
			Email:    "gin@example.com",
			Password: "",
		}
		err := Validator.ValidatePayloads(payload)

		// Assertions
		assert.NotNil(t, err)
		assert.Equal(t, "Password: is a required field", err.Error())
	})

	t.Run("Test 2 | Detect invalid format", func(t *testing.T) {
		payload := DummyPayload{
			Email:    "this.is-not_an.email",
			Password: "password876@",
		}
		err := Validator.ValidatePayloads(payload)

		// Assertions
		assert.NotNil(t, err)
		assert.Equal(t, "Email: is not a valid email address", err.Error())
	})

	t.Run("Test 3 | Error with params", func(t *testing.T) {
		payload := DummyPayload{
			Email:    "gin@example.com",
			Password: "short",
		}
		err := Validator.ValidatePayloads(payload)

		// Assertions
		assert.NotNil(t, err)
		assert.Equal(t, "Password: must be at least 6 characters long", err.Error())
	})

	t.Run("Test 4 | Validation passed", func(t *testing.T) {
		payload := DummyPayload{
			Email:    "gin@example.com",
			Password: "Password999!",
		}
		err := Validator.ValidatePayloads(payload)

		// Assertions
		assert.Nil(t, err)
	})
}
