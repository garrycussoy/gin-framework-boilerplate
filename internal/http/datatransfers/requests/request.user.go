package requests

import (
	Domains "gin-framework-boilerplate/internal/business/domains"
)

// Request-related variables in hanlder
type GetUserByEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type GetUsersRequest struct {
	BranchId string `json:"branch_id"`
	Start    string `json:"start"`
	End      string `json:"end"`
}

// Mapper which will be used to map request-related variables between handler and domain
func (req *GetUserByEmailRequest) ToUserDomain() *Domains.UserDomain {
	return &Domains.UserDomain{
		Email: req.Email,
	}
}

func (req *GetUsersRequest) ToUserFilterDomain() *Domains.UserFilterDomain {
	return &Domains.UserFilterDomain{
		BranchId: &req.BranchId,
		Start:    &req.Start,
		End:      &req.End,
	}
}
