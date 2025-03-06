package middlewares_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gin-framework-boilerplate/internal/config"
	"gin-framework-boilerplate/internal/http/middlewares"
	"gin-framework-boilerplate/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	jwtService          jwt.JWTService
	s                   *gin.Engine
	authBasicMiddleware gin.HandlerFunc
)

const (
	forEveryone = "/everyone"
)

func authenticatedHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "success",
	})
}

func setup() {
	// Load configuration file
	config.InitializeAppConfig(true)

	jwtService = jwt.NewJWTService(config.AppConfig.JWTSecret, config.AppConfig.JWTIssuer, config.AppConfig.JWTExpired)
	authBasicMiddleware = middlewares.NewAuthMiddleware(jwtService)

	s = gin.New()
	s.GET(forEveryone, authBasicMiddleware, authenticatedHandler)
}

func generateToken() (token string, err error) {
	token, err = jwtService.GenerateToken("ddfcea5c-d919-4a8f-a631-4ace39337s3a", "Admin", "gin@example.com")
	return
}

func getBasicToken() (string, error) {
	return generateToken()
}

func TestAuthMiddleware(t *testing.T) {
	setup()

	t.Run("Test 1 | User authenticated successfully", func(t *testing.T) {
		token, err := getBasicToken()
		if err != nil {
			t.Error(err)
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, forEveryone, nil)

		r.Header.Set("Content-Type", "application/json")
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		// Perform request
		s.ServeHTTP(w, r)
		body := w.Body.String()

		// Assertions
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Contains(t, w.Result().Header.Get("Content-Type"), "application/json")
		assert.Contains(t, body, "success")
	})

	t.Run("Test 2 | Invalid token", func(t *testing.T) {
		token := "invalid-token"

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, forEveryone, nil)

		r.Header.Set("Content-Type", "application/json")
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		// Perform request
		s.ServeHTTP(w, r)
		body := w.Body.String()

		// Assertions
		assert.Equal(t, http.StatusUnauthorized, w.Result().StatusCode)
		assert.Contains(t, w.Result().Header.Get("Content-Type"), "application/json")
		assert.Contains(t, body, "Invalid token")
	})

	t.Run("Test 3 | Must contain Bearer", func(t *testing.T) {
		token, err := getBasicToken()
		if err != nil {
			t.Error(err)
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, forEveryone, nil)

		r.Header.Set("Content-Type", "application/json")
		r.Header.Set("Authorization", fmt.Sprintf("Token %s", token))

		// Perform request
		s.ServeHTTP(w, r)
		body := w.Body.String()

		// Assertions
		assert.Equal(t, http.StatusUnauthorized, w.Result().StatusCode)
		assert.Contains(t, w.Result().Header.Get("Content-Type"), "application/json")
		assert.Contains(t, body, "Token must contain Bearer")
	})

	t.Run("Test 4 | Invalid format", func(t *testing.T) {
		token, err := getBasicToken()
		if err != nil {
			t.Error(err)
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, forEveryone, nil)

		r.Header.Set("Content-Type", "application/json")
		r.Header.Set("Authorization", fmt.Sprintf("Bearer token: %s", token))

		// Perform request
		s.ServeHTTP(w, r)
		body := w.Body.String()

		// Assertions
		assert.Equal(t, http.StatusUnauthorized, w.Result().StatusCode)
		assert.Contains(t, w.Result().Header.Get("Content-Type"), "application/json")
		assert.Contains(t, body, "Invalid header format")
	})
}
