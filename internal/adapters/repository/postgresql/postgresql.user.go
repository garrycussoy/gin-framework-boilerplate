package postgresql

import (
	"context"

	Repository "gin-framework-boilerplate/internal/ports/repository"
	"gin-framework-boilerplate/internal/ports/repository/records"

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

// Function to get registered user data by email
func (r *postgresqlUserRepository) GetUserByEmail(ctx context.Context, email string) (user records.User, err error) {
	// Get related user based on email
	err = r.conn.GetContext(ctx, &user, `SELECT * FROM "user" WHERE "email" = $1`, email)
	if err != nil {
		return records.User{}, err
	}

	return user, nil
}
