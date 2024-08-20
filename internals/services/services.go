package services

import (
	"github.com/Pr3c10us/boilerplate/internals/infrastructures/adapters"
	"github.com/Pr3c10us/boilerplate/internals/services/authentication"
)

type Services struct {
	AuthenticationServices authentication.Services
}

func NewServices(adapters *adapters.Adapters) *Services {
	return &Services{
		AuthenticationServices: authentication.NewAuthenticationServices(adapters.AuthenticationRepository),
	}
}
