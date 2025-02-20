package requests

import (
	Domains "gin-framework-boilerplate/internal/business/domains"
)

// Request-related variables in hanlder
type GetUserByEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// Mapper which will be used to map request-related variables between handler and domain
func (req *GetUserByEmailRequest) ToUserDomain() *Domains.UserDomain {
	return &Domains.UserDomain{
		Email: req.Email,
	}
}
