package authentication

import (
	"database/sql"
	"github.com/Pr3c10us/boilerplate/internals/domains/authentication"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthenticationRepositoryPG(db *sql.DB) authentication.Repository {
	return &AuthRepository{
		db,
	}
}
