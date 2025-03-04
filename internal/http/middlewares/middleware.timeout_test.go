package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"gin-framework-boilerplate/internal/config"
	"gin-framework-boilerplate/internal/http/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test TimeoutMiddleware functionality
func TestTimeoutMiddleware(t *testing.T) {
	// Load env variable
	config.InitializeAppConfig(true)

	// Initialize some variables
	router := gin.New()

	// Define router to test timeout middleware (passed)
	router.GET("/test-timeout", middlewares.TimeoutMiddleware(), func(c *gin.Context) {
		c.String(http.StatusOK, "Ok!")
	})

	// Define router to test timeout middleware (timeout exceeds)
	router.GET("/test-timeout-occur", middlewares.TimeoutMiddleware(), func(c *gin.Context) {
		time.Sleep(1500 * time.Millisecond)
		c.String(http.StatusOK, "Ok!")
	})

	t.Run("Test 1 | A function runs properly under given timeout", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test-timeout", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Test 2 | Timeout occurs", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test-timeout-occur", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusRequestTimeout, w.Code)
	})
}
