package usecases

import (
	"context"

	Domains "gin-framework-boilerplate/internal/business/domains"
	"gin-framework-boilerplate/internal/ports/repository"
	Errors "gin-framework-boilerplate/pkg/errors"
	"gin-framework-boilerplate/pkg/helpers"

	"gin-framework-boilerplate/pkg/jwt"
)

type authUsecase struct {
	jwtService jwt.JWTService
	userRepo   repository.UserRepository
}

func NewAuthUsecase(jwtService jwt.JWTService, userRepo repository.UserRepository) Domains.AuthUsecase {
	return &authUsecase{
		jwtService: jwtService,
		userRepo:   userRepo,
	}
}

func (authUC *authUsecase) UserLogin(ctx context.Context, inDom *Domains.UserLoginDomain) (Domains.UserLoginDomain, Errors.CustomError) {
	// Setup base response
	resp := Domains.UserLoginDomain{}
	var err error

	// Get user by email
	user, err := authUC.userRepo.GetUserByEmail(ctx, inDom.Email)
	if err != nil {
		return resp, Errors.AuthDomainError(400, "Invalid email or password")
	}

	// Validate password
	if !helpers.ValidateHash(inDom.Password, user.Password) {
		return resp, Errors.AuthDomainError(400, "Invalid email or password")
	}

	// Generate token
	resp.Token, err = authUC.jwtService.GenerateToken(user.Id, "Super Developer", user.Email)
	if err != nil {
		return resp, Errors.AuthDomainError(500, err.Error())
	}

	return resp, nil
}
