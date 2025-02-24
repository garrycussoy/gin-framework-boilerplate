package domains

import (
	"context"
	"time"

	DTO "gin-framework-boilerplate/internal/ports/repository/dto"
	Records "gin-framework-boilerplate/internal/ports/repository/records"
	Errors "gin-framework-boilerplate/pkg/errors"
	Helpers "gin-framework-boilerplate/pkg/helpers"
)

// User-related variables which will be used accross domain
type UserDomain struct {
	Id          string
	FullName    string
	Email       string
	PhoneNumber string
	Password    string
	BranchId    *string
	CreatedAt   time.Time
	CreatedBy   *string
	UpdatedAt   *time.Time
	UpdatedBy   *string
	Custom1     Helpers.JSONB
	Custom2     Helpers.JSONB
	Custom3     Helpers.JSONB
}

type UserFilterDomain struct {
	BranchId *string
	Start    *string
	End      *string
}

// User-related mapper which will be used accross domain
func FromUserToUserDomain(rec Records.User) UserDomain {
	return UserDomain{
		Id:          rec.Id,
		FullName:    rec.FullName,
		Email:       rec.Email,
		PhoneNumber: rec.PhoneNumber,
		Password:    rec.Password,
		BranchId:    rec.BranchId,
		CreatedAt:   rec.CreatedAt,
		CreatedBy:   rec.CreatedBy,
		UpdatedAt:   rec.UpdatedAt,
		UpdatedBy:   rec.UpdatedBy,
		Custom1:     rec.Custom1,
		Custom2:     rec.Custom2,
		Custom3:     rec.Custom3,
	}
}

func FromUserArrayToUserDomainArray(recArr []Records.User) []UserDomain {
	var domArr []UserDomain
	for _, rec := range recArr {
		domArr = append(domArr, FromUserToUserDomain(rec))
	}
	return domArr
}

func (dom *UserFilterDomain) FromUserFilterDomainToUserFilterDTO() DTO.UserFilterDto {
	return DTO.UserFilterDto{
		BranchId: dom.BranchId,
		Start:    dom.Start,
		End:      dom.End,
	}
}

// Interface for User domain
type UserUsecase interface {
	GetUsers(ctx context.Context, inDom *UserFilterDomain) ([]UserDomain, Errors.CustomError)
	GetUserByEmail(ctx context.Context, inDom *UserDomain) (UserDomain, Errors.CustomError)
}
