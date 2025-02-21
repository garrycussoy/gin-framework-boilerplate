package usecases

import (
	"context"

	Domains "gin-framework-boilerplate/internal/business/domains"
	"gin-framework-boilerplate/internal/ports/repository"
	Errors "gin-framework-boilerplate/pkg/errors"
)

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) Domains.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
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

	// Get user by email
	user, err := userUC.userRepo.GetUserByEmail(ctx, inDom.Email)
	if err != nil {
		return resp, Errors.AuthDomainError(400, "Email is not registered")
	}
	resp = Domains.FromUserToUserDomain(user)

	return resp, nil
}
