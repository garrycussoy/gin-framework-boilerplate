package jwt_test

import (
	"fmt"
	"testing"
	"time"

	"gin-framework-boilerplate/internal/config"
	"gin-framework-boilerplate/pkg/jwt"

	"github.com/stretchr/testify/assert"
)

// Test GenerateToken function
func TestGenerateToken(t *testing.T) {
	jwtService := jwt.NewJWTService(config.AppConfig.JWTSecret, config.AppConfig.JWTIssuer, config.AppConfig.JWTExpired)
	token, err := jwtService.GenerateToken("d2bd47f6-892e-417f-8a82-22d6db743f5f", "Admin", "john.doe@example.com")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

// Test ParseToken function
func TestParseToken(t *testing.T) {
	t.Run("With Valid Token", func(t *testing.T) {
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

	t.Run("With Invalid Token", func(t *testing.T) {
		jwtService := jwt.NewJWTService(config.AppConfig.JWTSecret, config.AppConfig.JWTIssuer, config.AppConfig.JWTExpired)

		_, err := jwtService.ParseToken("invalid_token")
		assert.Error(t, err)
		assert.Equal(t, "token is not valid", err.Error())
	})
}
