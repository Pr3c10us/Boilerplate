package command

import (
	"github.com/Pr3c10us/boilerplate/internals/domains/authentication"
)

type UpdateRefreshVersion interface {
	Handle(user *authentication.GetUserParams) (*authentication.User, error)
}

type updateRefreshVersion struct {
	authenticationRepository authentication.Repository
}

func (service *updateRefreshVersion) Handle(params *authentication.GetUserParams) (*authentication.User, error) {
	return service.authenticationRepository.UpdateUser(&authentication.User{ID: params.ID})
}

func NewUpdateRefreshVersion(repository authentication.Repository) UpdateRefreshVersion {
	return &updateRefreshVersion{
		repository,
	}
}
