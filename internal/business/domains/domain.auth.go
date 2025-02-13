package domains

import (
	"context"

	Errors "gin-framework-boilerplate/pkg/errors"
)

// Auth-related variables which will be used accross domain
type UserLoginDomain struct {
	Email    string
	Password string
	Token    string
}

// Interface for Auth domain
type AuthUsecase interface {
	UserLogin(ctx context.Context, inDom *UserLoginDomain) (UserLoginDomain, Errors.CustomError)
}
