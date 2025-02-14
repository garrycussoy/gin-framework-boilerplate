package records

import (
	"time"
	// "gin-framework-boilerplate/pkg/helpers"
)

type User struct {
	Id          string      `db:"id"`
	FullName    string      `db:"full_name"`
	Email       string      `db:"email"`
	PhoneNumber string      `db:"phone_number"`
	Password    string      `db:"password"`
	BranchId    *string     `db:"branch_id"`
	CreatedAt   time.Time   `db:"created_at"`
	CreatedBy   *string     `db:"created_by"`
	UpdatedAt   *time.Time  `db:"updated_at"`
	UpdatedBy   *string     `db:"updated_by"`
	Custom1     interface{} `db:"custom_1"`
	Custom2     interface{} `db:"custom_2"`
	Custom3     interface{} `db:"custom_3"`
}
