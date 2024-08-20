package authentication

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID                  uuid.UUID `json:"id"`
	Email               string    `json:"email"`
	FirstName           string    `json:"firstName"`
	LastName            string    `json:"lastName"`
	FullName            string    `json:"fullName"`
	AvatarURL           string    `json:"avatarURL"`
	Location            string    `json:"location"`
	RefreshTokenVersion int       `json:"refreshTokenVersion"`
	CreatedAt           time.Time `json:"createAt"`
}

type GetUserParams struct {
	Email string    `json:"email"    binding:"omitempty,email"`
	ID    uuid.UUID `json:"id"       binding:"omitempty,uuid"`
}
