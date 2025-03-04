package middlewares_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gin-framework-boilerplate/internal/config"
	"gin-framework-boilerplate/internal/constants"
	"gin-framework-boilerplate/internal/http/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test CORSMiddleware functionality
func TestCORSMiddleware(t *testing.T) {
	router := gin.New()
	router.Use(middlewares.CORSMiddleware())

	router.GET("/cors", func(c *gin.Context) {
		fmt.Println("This is header", c.Request.Header)
		c.String(http.StatusOK, "Ok!")
	})

	t.Run("Test 1 | Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/cors", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, config.AppConfig.AllowedCORS, w.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, constants.AllowedMethods, w.Header().Get("Access-Control-Allow-Methods"))
		assert.Equal(t, constants.AllowedHeader, w.Header().Get("Access-Control-Allow-Headers"))
		assert.Equal(t, "true", w.Header().Get("Access-Control-Allow-Credentials"))
	})

	t.Run("Test 2 | Method not allowed", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("CUSTOM", "/test", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusForbidden, w.Code)
		assert.Equal(t, config.AppConfig.AllowedCORS, w.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, constants.AllowedMethods, w.Header().Get("Access-Control-Allow-Methods"))
		assert.Equal(t, constants.AllowedHeader, w.Header().Get("Access-Control-Allow-Headers"))
		assert.Equal(t, "true", w.Header().Get("Access-Control-Allow-Credentials"))
	})

	t.Run("Test 3 | Header not allowed", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("misc", "something")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusForbidden, w.Code)
		assert.Equal(t, config.AppConfig.AllowedCORS, w.Header().Get("Access-Control-Allow-Origin"))
		assert.Equal(t, constants.AllowedMethods, w.Header().Get("Access-Control-Allow-Methods"))
		assert.Equal(t, constants.AllowedHeader, w.Header().Get("Access-Control-Allow-Headers"))
		assert.Equal(t, "true", w.Header().Get("Access-Control-Allow-Credentials"))
	})
}
