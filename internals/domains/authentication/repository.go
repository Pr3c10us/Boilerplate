package authentication

import (
	"github.com/markbates/goth"
)

type Repository interface {
	AddUser(user *goth.User) (*User, error)
	GetUserDetails(params *GetUserParams) (*User, error)
	UpdateUser(user *User) (*User, error)
}
