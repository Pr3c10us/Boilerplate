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

func TestAuthRepository_UpdateUser(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		params *authentication.User
	}
	db, dbMock := mock.NewMockDB()
	id, email := uuid.New(), "mr@man.strong"
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        *authentication.User
		expectedErr error
	}{
		{
			name: "success",
			fields: fields{
				db: db,
			},
			args: args{
				params: &authentication.User{
					ID:        uuid.New(),
					Email:     email,
					FirstName: "Mr",
					LastName:  "Man",
					AvatarURL: "",
					Location:  "",
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
			profileMap := map[string]interface{}{
				"refresh_token_version": sq.Expr("refresh_token_version + ?"),
			}
			query, _, err := sq.
				Update("users").
				SetMap(profileMap).Where(sq.Eq{"id": "?"}).
				Suffix(`RETURNING 
							"id",
							"email",
							"first_name",
							"last_name",
							"full_name",
							"avatar_url",
							"location",
							"created_at"`).
				PlaceholderFormat(sq.Dollar).
				ToSql()
			if err != nil {
				t.Error(err)
			}
			regexQuery := regexp.QuoteMeta(query)
			prep := dbMock.ExpectPrepare(regexQuery)
			if tt.expectedErr != nil {
				prep.ExpectQuery().WithArgs(1, tt.args.params.ID).WillReturnError(tt.expectedErr)
			} else {
				rows := sqlmock.NewRows([]string{"id", "email", "first_name", "last_name", "full_name", "avatar_url", " location", " created_at"}).
					AddRow(tt.want.ID.String(), tt.want.Email, tt.want.FirstName, tt.want.LastName, tt.want.FullName, tt.want.AvatarURL, tt.want.Location, tt.want.CreatedAt)
				prep.ExpectQuery().WithArgs(1, tt.args.params.ID).WillReturnRows(rows)
			}

			got, adapterErr := repo.UpdateUser(tt.args.params)
			assert.Equalf(t, tt.want, got, "UpdateUserDetails(%v)", tt.args.params)
			assert.Equal(t, tt.expectedErr, adapterErr)
		})
	}
}
