package usecases

import (
	"context"

	Domains "gin-framework-boilerplate/internal/business/domains"
	ESBPorts "gin-framework-boilerplate/internal/ports/clients/esb"
	Repository "gin-framework-boilerplate/internal/ports/repository"
	Errors "gin-framework-boilerplate/pkg/errors"
)

type userUsecase struct {
	userRepo  Repository.UserRepository
	esbClient ESBPorts.ESBClient
}

func NewUserUsecase(userRepo Repository.UserRepository, esbClient ESBPorts.ESBClient) Domains.UserUsecase {
	return &userUsecase{
		userRepo:  userRepo,
		esbClient: esbClient,
	}
}

func (userUC *userUsecase) GetUsers(ctx context.Context, inDom *Domains.UserFilterDomain) ([]Domains.UserDomain, Errors.CustomError) {
	// Setup base response
	resp := []Domains.UserDomain{}
	var err error

	// Get users based on given filter
	users, err := userUC.userRepo.GetUsers(ctx, inDom.FromUserFilterDomainToUserFilterDTO())
	if err != nil {
		return resp, Errors.UserRepositoryError(500, err.Error())
	}
	resp = Domains.FromUserArrayToUserDomainArray(users)

	return resp, nil
}

func (userUC *userUsecase) GetUserByEmail(ctx context.Context, inDom *Domains.UserDomain) (Domains.UserDomain, Errors.CustomError) {
	// Setup base response
	resp := Domains.UserDomain{}
	var err error

	// Sample request to ESB client
	// NOTES : Following code can be removed later as this is just an example
	userUC.esbClient.Sample(ctx)

	// Get user by email
	user, err := userUC.userRepo.GetUserByEmail(ctx, inDom.Email)
	if err != nil {
		return resp, Errors.AuthDomainError(400, "Email is not registered")
	}
	resp = Domains.FromUserToUserDomain(user)

	return resp, nil
}
