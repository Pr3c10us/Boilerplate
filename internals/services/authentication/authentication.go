package authentication

import (
	"github.com/Pr3c10us/boilerplate/internals/domains/authentication"
	"github.com/Pr3c10us/boilerplate/internals/services/authentication/command"
	"github.com/Pr3c10us/boilerplate/internals/services/authentication/queries"
)

type Services struct {
	Commands
	Queries
}

type Commands struct {
	Authenticate         command.Authenticate
	UpdateRefreshVersion command.UpdateRefreshVersion
}

type Queries struct {
	GetUserDetails queries.GetUserDetails
}

func NewAuthenticationServices(authenticationRepository authentication.Repository) Services {
	return Services{
		Commands: Commands{
			Authenticate:         command.NewAuthenticate(authenticationRepository),
			UpdateRefreshVersion: command.NewUpdateRefreshVersion(authenticationRepository),
		},
		Queries: Queries{
			GetUserDetails: queries.NewGetUserDetails(authenticationRepository),
		},
	}
}
