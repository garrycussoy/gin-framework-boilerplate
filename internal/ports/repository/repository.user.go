package repository

import (
	"context"

	DTO "gin-framework-boilerplate/internal/ports/repository/dto"
	Records "gin-framework-boilerplate/internal/ports/repository/records"
)

type UserRepository interface {
	GetUsers(ctx context.Context, userFilter DTO.UserFilterDto) (users []Records.User, err error)
	GetUserByEmail(ctx context.Context, email string) (user Records.User, err error)
}
