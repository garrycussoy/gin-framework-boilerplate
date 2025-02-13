package responses

import (
	Domains "gin-framework-boilerplate/internal/business/domains"
)

// Response-related variables in hanlder
type UserLoginResponse struct {
	Token string `json:"token"`
}

// Mapper which will be used to map response-related variables between handler and domain
func FromUserLoginDomainToUserLoginResponse(dom Domains.UserLoginDomain) UserLoginResponse {
	return UserLoginResponse{
		Token: dom.Token,
	}
}
