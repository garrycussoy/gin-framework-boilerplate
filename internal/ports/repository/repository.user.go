package repository

import (
	"context"

	"gin-framework-boilerplate/internal/ports/repository/records"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (user records.User, err error)
}
