package handlers

import (
	Domains "gin-framework-boilerplate/internal/business/domains"
	"gin-framework-boilerplate/internal/http/datatransfers/requests"
	"gin-framework-boilerplate/internal/http/datatransfers/responses"
	Errors "gin-framework-boilerplate/pkg/errors"

	"gin-framework-boilerplate/pkg/validators"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase Domains.UserUsecase
}

func NewUserHandler(usecase Domains.UserUsecase) UserHandler {
	return UserHandler{
		usecase: usecase,
	}
}

// @Summary Retrieve user data based on email
// @Description System will receive an email input, then find related user data in the database
// @Tags User
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param email path string true "Email from a user whom we want to collect the data"
// @Success 200 {object} responses.GetUserByEmailResponse
// @Router /user/{email} [get]
func (userH UserHandler) GetUserByEmail(ctx *gin.Context) {
	// Extract query params
	GetUserByEmailRequest := requests.GetUserByEmailRequest{
		Email: ctx.Param("email"),
	}

	// Validate the params
	if err := validators.ValidatePayloads(GetUserByEmailRequest); err != nil {
		ErrorResponse(ctx, Errors.ValidationFailed(err.Error()))
		return
	}

	// The main process of retrieving user data based on given email
	resp, err := userH.usecase.GetUserByEmail(ctx.Request.Context(), GetUserByEmailRequest.ToUserDomain())
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	// Return response
	SuccessResponse(ctx, responses.FromUserDomainToGetUserByEmailResponse(resp))
}
