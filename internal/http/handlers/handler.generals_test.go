package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	Handlers "gin-framework-boilerplate/internal/http/handlers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	generalHandler Handlers.GeneralHandler
)

func generalTestSetup() {
	generalHandler = Handlers.NewGeneralHandler()

	// Create gin engine
	s = gin.Default()
}

// Test HealthCheck function
func TestHealthCheck(t *testing.T) {
	generalTestSetup()

	// Define route
	s.GET("/health-check", generalHandler.HealthCheck)
	t.Run("Test 1 | Success", func(t *testing.T) {
		// Do some setup
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/health-check", nil)
		r.Header.Set("Content-Type", "application/json")

		// Perform request
		s.ServeHTTP(w, r)

		var body Handlers.BaseResponse
		json.Unmarshal(w.Body.Bytes(), &body)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Contains(t, w.Result().Header.Get("Content-Type"), "application/json")
		assert.Contains(t, body.Message, "Success")
	})
}
