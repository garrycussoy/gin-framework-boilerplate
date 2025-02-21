package domains

import (
	"context"

	DTO "gin-framework-boilerplate/internal/ports/repository/dto"
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

type UserFilterDomain struct {
	BranchId *string
	Start    *string
	End      *string
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
