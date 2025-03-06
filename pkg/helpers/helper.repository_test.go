package helpers_test

import (
	"testing"

	"gin-framework-boilerplate/pkg/helpers"

	"github.com/stretchr/testify/assert"
)

func TestFilterQueryGenerator(t *testing.T) {
	t.Run("Test 1 | Success generating the query for all types", func(t *testing.T) {
		filter := map[string][]*string{
			"email": {helpers.CreatePointerString("gin@example.com")},
			"age":   {helpers.CreatePointerString("25"), helpers.CreatePointerString(">")},
			"start": {helpers.CreatePointerString("2025-01-01 00:00:00"), helpers.CreatePointerString(">"), helpers.CreatePointerString("created_at")},
		}
		query := helpers.FilterQueryGenerator("User", filter)

		// Assertions
		assert.Contains(t, query, `SELECT * FROM "User"`)
		assert.Contains(t, query, `"email" = 'gin@example.com'`)
		assert.Contains(t, query, `"age" > '25'`)
		assert.Contains(t, query, `"created_at" > '2025-01-01 00:00:00'`)
	})
}
