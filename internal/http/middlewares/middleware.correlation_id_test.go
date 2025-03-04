package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gin-framework-boilerplate/internal/http/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Initialize temp variable to store correlation ID
var correlationId any

// Test CorrelationIdMiddleware functionality
func TestCorrelationIdMiddleware(t *testing.T) {
	// Initialize some variables
	router := gin.New()

	// Define router to test correlation Id middleware (passed)
	router.GET("/test-corr-id", middlewares.CorrelationIdMiddleware(), func(c *gin.Context) {
		correlationId = c.Value("CorrelationId")
		c.String(http.StatusOK, "Ok!")
	})

	t.Run("Test 1 | Ensure correlation Id is generated uniquely", func(t *testing.T) {
		// First request
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test-corr-id", nil)
		router.ServeHTTP(w, req)
		corrId1 := correlationId

		// Second request
		req, _ = http.NewRequest("GET", "/test-corr-id", nil)
		router.ServeHTTP(w, req)
		corrId2 := correlationId

		// Assert that the correlation Id is generated uniquely
		assert.NotEmpty(t, corrId1)
		assert.NotEmpty(t, corrId2)
		assert.False(t, corrId1 == corrId2)
	})
}
