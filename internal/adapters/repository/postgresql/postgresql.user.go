package postgresql

import (
	"context"

	Repository "gin-framework-boilerplate/internal/ports/repository"
	DTO "gin-framework-boilerplate/internal/ports/repository/dto"
	Records "gin-framework-boilerplate/internal/ports/repository/records"
	"gin-framework-boilerplate/pkg/helpers"

	"github.com/jmoiron/sqlx"
)

type postgresqlUserRepository struct {
	conn *sqlx.DB
}

func NewUserRepository(conn *sqlx.DB) Repository.UserRepository {
	return &postgresqlUserRepository{
		conn: conn,
	}
}

// Function to get list of users based on given filter
func (r *postgresqlUserRepository) GetUsers(ctx context.Context, userFilter DTO.UserFilterDto) (users []Records.User, err error) {
	// Build the query
	filter := map[string][]*string{
		"branch_id": {userFilter.BranchId},
	}
	query := helpers.FilterQueryGenerator("user", filter)

	// Get related user based on email
	err = r.conn.GetContext(ctx, users, query)
	if err != nil {
		return []Records.User{}, err
	}

	return users, nil
}

// Function to get registered user data by email
func (r *postgresqlUserRepository) GetUserByEmail(ctx context.Context, email string) (user Records.User, err error) {
	// Build the query
	filter := map[string][]*string{
		"email": {&email},
	}
	query := helpers.FilterQueryGenerator("user", filter)

	// Get related user based on email
	err = r.conn.GetContext(ctx, &user, query)
	if err != nil {
		return Records.User{}, err
	}

	return user, nil
}
