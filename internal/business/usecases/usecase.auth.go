package usecases

import (
	"context"

	Domains "gin-framework-boilerplate/internal/business/domains"
	Errors "gin-framework-boilerplate/pkg/errors"
	"gin-framework-boilerplate/pkg/helpers"
)

type authUsecase struct {
	// patientRepo    Domains.PatientRepository
	// masterDataRepo Domains.MasterDataRepository
}

func NewAuthUsecase() Domains.AuthUsecase {
	return &authUsecase{}
}

func (authUC *authUsecase) UserLogin(ctx context.Context, inDom *Domains.UserLoginDomain) (Domains.UserLoginDomain, Errors.CustomError) {
	// Validate password
	if !helpers.ValidateHash(inDom.Password, "Password") {
		return Domains.UserLoginDomain{}, Errors.AuthDomainError(400, "Invalid email or password")
	}

	return Domains.UserLoginDomain{}, nil
}
