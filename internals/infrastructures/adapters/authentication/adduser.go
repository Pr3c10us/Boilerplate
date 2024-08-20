package authentication

import (
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/Pr3c10us/boilerplate/internals/domains/authentication"
	"github.com/Pr3c10us/boilerplate/packages/appError"
	"github.com/markbates/goth"
)

func (repo *AuthRepository) AddUser(user *goth.User) (*authentication.User, error) {
	query, _, err := sq.Insert("users").
		Columns("email", "first_name", "last_name", "avatar_url", "location").
		Values("?", "?", "?", "?", "?").
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
		return nil, err
	}
	statement, err := repo.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	var newUser authentication.User
	switch err = statement.QueryRow(
		user.Email,
		user.FirstName,
		user.LastName,
		user.AvatarURL,
		user.Location,
	).Scan(
		&newUser.ID,
		&newUser.Email,
		&newUser.FirstName,
		&newUser.LastName,
		&newUser.FullName,
		&newUser.AvatarURL,
		&newUser.Location,
		&newUser.CreatedAt,
	); {
	case errors.Is(err, sql.ErrNoRows):
		var userNotCreatedErr = errors.New("failed to create user")
		return nil, appError.InternalServerError(userNotCreatedErr)
	case err == nil:
		return &newUser, nil
	default:
		return nil, err
	}
}
