package usecases

import (
	"context"

	Domains "gin-framework-boilerplate/internal/business/domains"
	Errors "gin-framework-boilerplate/pkg/errors"
)

type authUsecase struct {
	// patientRepo    Domains.PatientRepository
	// masterDataRepo Domains.MasterDataRepository
}

func NewAuthUsecase() Domains.AuthUsecase {
	return &authUsecase{}
}

func (authUC *authUsecase) UserLogin(ctx context.Context, inDom *Domains.UserLoginDomain) (Domains.UserLoginDomain, Errors.CustomError) {
	return Domains.UserLoginDomain{}, nil
}
