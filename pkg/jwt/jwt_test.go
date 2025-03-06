package jwt_test

import (
	"fmt"
	"testing"
	"time"

	"gin-framework-boilerplate/internal/config"
	"gin-framework-boilerplate/pkg/jwt"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	t.Run("Test 1 | Success generating token", func(t *testing.T) {
		jwtService := jwt.NewJWTService(config.AppConfig.JWTSecret, config.AppConfig.JWTIssuer, config.AppConfig.JWTExpired)
		token, err := jwtService.GenerateToken("d2bd47f6-892e-417f-8a82-22d6db743f5f", "Admin", "john.doe@example.com")
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})
}

func TestParseToken(t *testing.T) {
	t.Run("Test 1 | With valid token", func(t *testing.T) {
		jwtService := jwt.NewJWTService(config.AppConfig.JWTSecret, config.AppConfig.JWTIssuer, config.AppConfig.JWTExpired)
		config.AppConfig.JWTExpired = 5

		token, _ := jwtService.GenerateToken("d2bd47f6-892e-417f-8a82-22d6db743f5f", "Admin", "john.doe@example.com")

		claims, err := jwtService.ParseToken(token)
		fmt.Println("This is an expire token!", claims.StandardClaims.ExpiresAt)
		assert.NoError(t, err)
		assert.Equal(t, "d2bd47f6-892e-417f-8a82-22d6db743f5f", claims.UserId)
		assert.Equal(t, "john.doe@example.com", claims.Email)
		assert.True(t, claims.StandardClaims.ExpiresAt >= time.Now().UTC().Unix())
		assert.Equal(t, config.AppConfig.JWTIssuer, claims.StandardClaims.Issuer)
		assert.True(t, claims.StandardClaims.IssuedAt <= time.Now().UTC().Unix())
	})

	t.Run("Test 2 | With invalid token", func(t *testing.T) {
		jwtService := jwt.NewJWTService(config.AppConfig.JWTSecret, config.AppConfig.JWTIssuer, config.AppConfig.JWTExpired)

		_, err := jwtService.ParseToken("invalid_token")
		assert.Error(t, err)
		assert.Equal(t, "token is not valid", err.Error())
	})
}
