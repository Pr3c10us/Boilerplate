package command

import (
	authentication2 "github.com/Pr3c10us/boilerplate/internals/domains/authentication"
	"github.com/Pr3c10us/boilerplate/internals/infrastructures/adapters/authentication"
	"github.com/google/uuid"
	"github.com/markbates/goth"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_authenticate_Handle(t *testing.T) {
	id, email := uuid.New(), "mr@man.strong"
	type args struct {
		user *goth.User
	}
	tests := []struct {
		name          string
		args          args
		want          *authentication2.User
		expectedError error
	}{
		{
			name: "user does not exist",
			args: args{
				user: &goth.User{
					Email:     email,
					FirstName: "Mr",
					LastName:  "Man",
					AvatarURL: "",
					Location:  "",
				},
			},
			want:          nil,
			expectedError: nil,
		},
		{
			name: "user exist",
			args: args{
				user: &goth.User{
					Email:     email,
					FirstName: "Mr",
					LastName:  "Man",
					AvatarURL: "",
					Location:  "",
				},
			},
			want: &authentication2.User{
				ID:        id,
				Email:     email,
				FirstName: "Mr",
				LastName:  "Man",
				FullName:  "Mr Man",
				AvatarURL: "",
				Location:  "",
				CreatedAt: time.Now(),
			},
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAuthRepository := new(authentication.MockRepository)
			mockAuthRepository.On(
				"GetUserDetails",
				&authentication2.GetUserParams{
					Email: tt.args.user.Email,
					ID:    uuid.Nil,
				}).
				Return(tt.want, tt.expectedError)
			mockAuthRepository.On("AddUser", tt.args.user).Return(tt.want, tt.expectedError)

			services := &authenticate{
				authenticationRepository: mockAuthRepository,
			}

			user, err := services.Handle(tt.args.user)
			assert.Equal(t, tt.want, user)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
