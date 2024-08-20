package authentication

import (
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/Pr3c10us/boilerplate/internals/domains/authentication"
	"github.com/Pr3c10us/boilerplate/packages/appError"
)

func (repo *AuthRepository) GetUserDetails(params *authentication.GetUserParams) (*authentication.User, error) {
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
		return nil, err
	}

	var statement *sql.Stmt
	statement, err = repo.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	var user authentication.User
	switch err = statement.QueryRow(params.Email, params.ID).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.FullName,
		&user.AvatarURL,
		&user.Location,
		&user.RefreshTokenVersion,
		&user.CreatedAt,
	); {
	case errors.Is(err, sql.ErrNoRows):
		var userNotFoundErr = errors.New("user does not exit")
		return nil, appError.NotFound(userNotFoundErr)
	case err == nil:
		return &user, nil
	default:
		return nil, err
	}
}
