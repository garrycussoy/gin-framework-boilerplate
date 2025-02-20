package responses

import (
	Domains "gin-framework-boilerplate/internal/business/domains"
)

// Response-related variables in hanlder
type GetUserByEmailResponse struct {
	Id       string `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

// Mapper which will be used to map response-related variables between handler and domain
func FromUserDomainToGetUserByEmailResponse(dom Domains.UserDomain) GetUserByEmailResponse {
	return GetUserByEmailResponse{
		Id:       dom.Id,
		FullName: dom.FullName,
		Email:    dom.Email,
	}
}
