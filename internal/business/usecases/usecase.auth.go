package usecases

import (
	"context"

	Domains "gin-framework-boilerplate/internal/business/domains"
	Errors "gin-framework-boilerplate/pkg/errors"
	"gin-framework-boilerplate/pkg/helpers"

	"gin-framework-boilerplate/pkg/jwt"
)

type authUsecase struct {
	jwtService jwt.JWTService
	// patientRepo    Domains.PatientRepository
	// masterDataRepo Domains.MasterDataRepository
}

func NewAuthUsecase(jwtService jwt.JWTService) Domains.AuthUsecase {
	return &authUsecase{
		jwtService: jwtService,
	}
}

func (authUC *authUsecase) UserLogin(ctx context.Context, inDom *Domains.UserLoginDomain) (Domains.UserLoginDomain, Errors.CustomError) {
	// Setup base response
	resp := Domains.UserLoginDomain{}
	var err error

	// Validate password
	if !helpers.ValidateHash(inDom.Password, "$2a$10$ZQ1qwpA7WEyAG6niUSHbjOlXmFB6F2N3GHT.mlhKY1.nsI8aep4kW") {
		return resp, Errors.AuthDomainError(400, "Invalid email or password")
	}

	// Generate token
	resp.Token, err = authUC.jwtService.GenerateToken("id-007", "Super Developer", inDom.Email)
	if err != nil {
		return resp, Errors.AuthDomainError(500, err.Error())
	}

	return resp, nil
}
