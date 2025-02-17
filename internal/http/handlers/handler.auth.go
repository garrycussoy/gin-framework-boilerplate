package handlers

import (
	Domains "gin-framework-boilerplate/internal/business/domains"
	"gin-framework-boilerplate/internal/http/datatransfers/requests"
	"gin-framework-boilerplate/internal/http/datatransfers/responses"
	Errors "gin-framework-boilerplate/pkg/errors"

	"gin-framework-boilerplate/pkg/validators"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	usecase Domains.AuthUsecase
}

func NewAuthHandler(usecase Domains.AuthUsecase) AuthHandler {
	return AuthHandler{
		usecase: usecase,
	}
}

// Handler to handle user login process
func (authH AuthHandler) UserLogin(ctx *gin.Context) {
	// Extract body request
	var UserLoginRequest requests.UserLoginRequest
	if err := ctx.ShouldBindJSON(&UserLoginRequest); err != nil {
		ErrorResponse(ctx, Errors.ValidationFailed(err.Error()))
		return
	}

	// Validate body request
	if err := validators.ValidatePayloads(UserLoginRequest); err != nil {
		ErrorResponse(ctx, Errors.ValidationFailed(err.Error()))
		return
	}

	// Do login process
	resp, err := authH.usecase.UserLogin(ctx.Request.Context(), UserLoginRequest.ToUserLoginDomain())
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	// Return response
	SuccessResponse(ctx, responses.FromUserLoginDomainToUserLoginResponse(resp))
}
