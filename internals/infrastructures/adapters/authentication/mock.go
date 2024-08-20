package authentication

import (
	"github.com/Pr3c10us/boilerplate/internals/domains/authentication"
	"github.com/markbates/goth"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) AddUser(params *goth.User) (*authentication.User, error) {
	args := mock.Called(params)
	return args.Get(0).(*authentication.User), args.Error(1)
}

func (mock *MockRepository) UpdateUser(params *authentication.User) (*authentication.User, error) {
	args := mock.Called(params)
	return args.Get(0).(*authentication.User), args.Error(1)
}

func (mock *MockRepository) GetUserDetails(params *authentication.GetUserParams) (*authentication.User, error) {
	args := mock.Called(params)
	return args.Get(0).(*authentication.User), args.Error(1)
}
