package command

import (
	"github.com/Pr3c10us/boilerplate/internals/domains/authentication"
	authentication2 "github.com/Pr3c10us/boilerplate/internals/infrastructures/adapters/authentication"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_updateRefreshVersion_Handle(t *testing.T) {
	id, email := uuid.New(), "mr@man.strong"
	type args struct {
		params *authentication.GetUserParams
	}
	tests := []struct {
		name          string
		args          args
		want          *authentication.User
		expectedError error
	}{
		{
			name: "success",
			args: args{
				params: &authentication.GetUserParams{
					ID: id,
				},
			},
			want: &authentication.User{
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
			mockAuthRepository := new(authentication2.MockRepository)
			mockAuthRepository.On("UpdateUser", &authentication.User{ID: tt.args.params.ID}).Return(tt.want, tt.expectedError)

			services := &updateRefreshVersion{
				authenticationRepository: mockAuthRepository,
			}

			user, err := services.Handle(tt.args.params)
			assert.Equal(t, tt.want, user)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
