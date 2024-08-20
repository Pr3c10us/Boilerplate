package authentication

import (
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/Pr3c10us/boilerplate/internals/domains/authentication"
	"github.com/Pr3c10us/boilerplate/packages/appError"
)

func (repo *AuthRepository) UpdateUser(user *authentication.User) (*authentication.User, error) {
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
		return nil, err
	}

	statement, err := repo.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	var newUser authentication.User
	switch err = statement.QueryRow(
		1,
		user.ID,
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
		var userNotUpdatedErr = errors.New("failed to update user")
		return nil, appError.NotFound(userNotUpdatedErr)
	case err == nil:
		return &newUser, nil
	default:
		return nil, err
	}
}
