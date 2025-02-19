package middlewares

import (
	"strings"

	"gin-framework-boilerplate/internal/http/handlers"
	Errors "gin-framework-boilerplate/pkg/errors"
	"gin-framework-boilerplate/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	jwtService jwt.JWTService
}

func NewAuthMiddleware(jwtService jwt.JWTService) gin.HandlerFunc {
	return (&AuthMiddleware{
		jwtService: jwtService,
	}).Handle
}

// Function which will be called everytime authentication is needed
func (m *AuthMiddleware) Handle(ctx *gin.Context) {
	// Validate that the request contains Authorization header
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		handlers.AbortResponse(ctx, Errors.AuthorizationFailed("Missing authorization header"))
		return
	}

	// Validate header format
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		handlers.AbortResponse(ctx, Errors.AuthorizationFailed("Invalid header format"))
		return
	}
	if headerParts[0] != "Bearer" {
		handlers.AbortResponse(ctx, Errors.AuthorizationFailed("Token must contain Bearer"))
		return
	}

	// Parse and validate JWT, then extract the claims
	_, err := m.jwtService.ParseToken(headerParts[1])
	if err != nil {
		handlers.AbortResponse(ctx, Errors.AuthorizationFailed("Invalid token"))
		return
	}

	ctx.Next()
}
