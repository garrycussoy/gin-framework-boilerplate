package domains

import (
	"context"

	Records "gin-framework-boilerplate/internal/ports/repository/records"
	Errors "gin-framework-boilerplate/pkg/errors"
)

// User-related variables which will be used accross domain
type UserDomain struct {
	Id       string
	Email    string
	Password string
	FullName string
}

// User-related mapper which will be used accross domain
func FromUserToUserDomain(rec Records.User) UserDomain {
	return UserDomain{
		Id:       rec.Id,
		Email:    rec.Email,
		Password: rec.Password,
		FullName: rec.FullName,
	}
}

// Interface for User domain
type UserUsecase interface {
	GetUserByEmail(ctx context.Context, inDom *UserDomain) (UserDomain, Errors.CustomError)
}
