package authentication

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	sq "github.com/Masterminds/squirrel"
	"github.com/Pr3c10us/boilerplate/internals/domains/authentication"
	"github.com/Pr3c10us/boilerplate/packages/mock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

func TestAuthenticationRepository_GetUserDetails(t *testing.T) {
	db, dbMock := mock.NewMockDB()
	id, email := uuid.New(), "mr@man.strong"
	type fields struct {
		db *sql.DB
	}
	type args struct {
		params *authentication.GetUserParams
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        *authentication.User
		expectedErr error
	}{
		{
			name: "success email",
			fields: fields{
				db: db,
			},
			args: args{
				params: &authentication.GetUserParams{
					Email: email,
					ID:    uuid.Nil,
				},
			},
			want: &authentication.User{
				ID:                  id,
				Email:               email,
				FirstName:           "Mr",
				LastName:            "Man",
				FullName:            "Mr Man",
				AvatarURL:           "",
				Location:            "",
				RefreshTokenVersion: 0,
				CreatedAt:           time.Now(),
			},
			expectedErr: nil,
		},
		{
			name: "success id",
			fields: fields{
				db: db,
			},
			args: args{
				params: &authentication.GetUserParams{
					Email: "",
					ID:    id,
				},
			},
			want: &authentication.User{
				ID:                  id,
				Email:               email,
				FirstName:           "Mr",
				LastName:            "Man",
				FullName:            "Mr Man",
				AvatarURL:           "",
				Location:            "",
				RefreshTokenVersion: 0,
				CreatedAt:           time.Now(),
			},
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := AuthRepository{
				db: tt.fields.db,
			}
			query, _, err := sq.Select(
				"id",
				"email",
				"COALESCE(first_name, '') AS first_name",
				"COALESCE(last_name, '') AS last_name",
				"COALESCE(full_name, '') AS full_name",
				"COALESCE(avatar_url, '') AS avatar_url",
				"COALESCE(location, '') AS location",
				"refresh_token_version",
				"created_at",
			).From("users").Where(sq.Or{sq.Eq{"email": "?"}, sq.Eq{"id": "?"}}).PlaceholderFormat(sq.Dollar).ToSql()
			if err != nil {
				t.Error(err)
			}
			regexQuery := regexp.QuoteMeta(query)
			prep := dbMock.ExpectPrepare(regexQuery)
			if tt.expectedErr != nil {
				prep.ExpectQuery().WithArgs(tt.args.params.Email, tt.args.params.ID).WillReturnError(tt.expectedErr)
			} else {
				rows := sqlmock.NewRows([]string{"id", "email", "first_name", "last_name", "full_name", "avatar_url", " location", "refresh_token_version", " created_at"}).
					AddRow(tt.want.ID.String(), tt.want.Email, tt.want.FirstName, tt.want.LastName, tt.want.FullName, tt.want.AvatarURL, tt.want.Location, tt.want.RefreshTokenVersion, tt.want.CreatedAt)
				prep.ExpectQuery().WithArgs(tt.args.params.Email, tt.args.params.ID).WillReturnRows(rows)
			}

			got, adapterErr := repo.GetUserDetails(tt.args.params)
			assert.Equalf(t, tt.want, got, "GetUserDetails(%v)", tt.args.params)
			assert.Equal(t, tt.expectedErr, adapterErr)
		})
	}
}
