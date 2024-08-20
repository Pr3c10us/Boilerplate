package queries

import "github.com/Pr3c10us/boilerplate/internals/domains/authentication"

type GetUserDetails interface {
	Handle(user *authentication.GetUserParams) (*authentication.User, error)
}

type getUserDetails struct {
	authenticationRepository authentication.Repository
}

func (service *getUserDetails) Handle(params *authentication.GetUserParams) (*authentication.User, error) {
	return service.authenticationRepository.GetUserDetails(params)
}

func NewGetUserDetails(repository authentication.Repository) GetUserDetails {
	return &getUserDetails{
		repository,
	}
}
