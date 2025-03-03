package helpers_test

import (
	"testing"

	"gin-framework-boilerplate/pkg/helpers"

	"github.com/stretchr/testify/assert"
)

// Test FilterQueryGenerator function
func TestFilterQueryGenerator(t *testing.T) {
	// Test 1 (Success)
	filter := map[string][]*string{
		"email": {helpers.CreatePointerString("gin@example.com")},
		"age":   {helpers.CreatePointerString("25"), helpers.CreatePointerString(">")},
		"start": {helpers.CreatePointerString("2025-01-01 00:00:00"), helpers.CreatePointerString(">"), helpers.CreatePointerString("created_at")},
	}
	query := helpers.FilterQueryGenerator("User", filter)

	// Assertions
	assert.Equal(t, `SELECT * FROM "User" WHERE "email" = 'gin@example.com' AND "age" > '25' AND "created_at" > '2025-01-01 00:00:00'`, query)
}
