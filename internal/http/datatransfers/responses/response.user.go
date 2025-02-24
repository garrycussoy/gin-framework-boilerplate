package responses

import (
	Domains "gin-framework-boilerplate/internal/business/domains"
	Helpers "gin-framework-boilerplate/pkg/helpers"
	"time"
)

// Response-related variables in hanlder
type UserResponse struct {
	Id          string        `json:"id"`
	FullName    string        `json:"full_name"`
	Email       string        `json:"email"`
	PhoneNumber string        `json:"phone_number"`
	Password    string        `json:"password"`
	BranchId    *string       `json:"branch_id"`
	CreatedAt   time.Time     `json:"created_at"`
	CreatedBy   *string       `json:"created_by"`
	UpdatedAt   *time.Time    `json:"updated_at"`
	UpdatedBy   *string       `json:"updated_by"`
	Custom1     Helpers.JSONB `json:"custom_1"`
	Custom2     Helpers.JSONB `json:"custom_2"`
	Custom3     Helpers.JSONB `json:"custom_3"`
}

// Mapper which will be used to map response-related variables between handler and domain
func FromUserDomainToUserResponse(dom Domains.UserDomain) UserResponse {
	return UserResponse{
		Id:          dom.Id,
		FullName:    dom.FullName,
		Email:       dom.Email,
		PhoneNumber: dom.PhoneNumber,
		Password:    dom.Password,
		BranchId:    dom.BranchId,
		CreatedAt:   dom.CreatedAt,
		CreatedBy:   dom.CreatedBy,
		UpdatedAt:   dom.UpdatedAt,
		UpdatedBy:   dom.UpdatedBy,
		Custom1:     dom.Custom1,
		Custom2:     dom.Custom2,
		Custom3:     dom.Custom3,
	}
}

func FromUserDomainArrayToUserResponseArray(domArr []Domains.UserDomain) []UserResponse {
	respArr := []UserResponse{}
	for _, dom := range domArr {
		respArr = append(respArr, FromUserDomainToUserResponse(dom))
	}
	return respArr
}
