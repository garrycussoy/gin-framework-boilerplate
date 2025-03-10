package requests

import (
	Domains "gin-framework-boilerplate/internal/business/domains"
)

// Request-related variables in hanlder
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Mapper which will be used to map request-related variables between handler and domain
func (req *UserLoginRequest) ToUserLoginDomain() *Domains.UserLoginDomain {
	return &Domains.UserLoginDomain{
		Email:    req.Email,
		Password: req.Password,
	}
}
