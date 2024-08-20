package command

import (
	"errors"
	"github.com/Pr3c10us/boilerplate/internals/domains/authentication"
	"github.com/Pr3c10us/boilerplate/packages/appError"
	"github.com/markbates/goth"
)

type Authenticate interface {
	Handle(user *goth.User) (*authentication.User, error)
}

type authenticate struct {
	authenticationRepository authentication.Repository
}

func (service *authenticate) Handle(user *goth.User) (*authentication.User, error) {
	var userNotFoundErr = errors.New("user does not exit")
	switch userFetched, err := service.authenticationRepository.GetUserDetails(
		&authentication.GetUserParams{Email: user.Email},
	); {
	case errors.Is(err, appError.NotFound(userNotFoundErr)) || userFetched == nil:
		userFetched, err = service.authenticationRepository.AddUser(user)
		if err != nil {
			return nil, err
		}
		return userFetched, nil
	case err == nil:
		return userFetched, nil
	default:
		return nil, err
	}
}

func NewAuthenticate(repository authentication.Repository) Authenticate {
	return &authenticate{
		repository,
	}
}
