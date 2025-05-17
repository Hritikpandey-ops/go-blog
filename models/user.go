package models

import (
	"database/sql"
	"time"

	"github.com/Hritikpandey-ops/blog-site/enums"
	"github.com/segmentio/ksuid"
)

type User struct {
	ID        ksuid.KSUID    `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"-"`
	Role      enums.UserRole `json:"role"`
	AvatarURL *string        `json:"avatar_url,omitempty"`
	Bio       *string        `json:"bio,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

// Save inserts the User into the database
func (u *User) Save(db *sql.DB) error {
	query := `
		INSERT INTO users (name, email, password, role, avatar_url, bio)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at;`

	err := db.QueryRow(
		query,
		u.Name,
		u.Email,
		u.Password,
		u.Role,
		u.AvatarURL,
		u.Bio,
	).Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)

	return err
}
