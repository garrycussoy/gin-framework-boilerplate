package responses

import (
	Domains "gin-framework-boilerplate/internal/business/domains"
)

// Response-related variables in hanlder
type UserResponse struct {
	Id       string `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

// Mapper which will be used to map response-related variables between handler and domain
func FromUserDomainToUserResponse(dom Domains.UserDomain) UserResponse {
	return UserResponse{
		Id:       dom.Id,
		FullName: dom.FullName,
		Email:    dom.Email,
	}
}

func FromUserDomainArrayToUserResponseArray(domArr []Domains.UserDomain) []UserResponse {
	respArr := []UserResponse{}
	for _, dom := range domArr {
		respArr = append(respArr, FromUserDomainToUserResponse(dom))
	}
	return respArr
}
