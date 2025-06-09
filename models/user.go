package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Hritikpandey-ops/blog-site/enums"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"-"`
	Role      enums.UserRole `json:"role"`
	AvatarURL *string        `json:"avatar_url,omitempty"`
	Bio       *string        `json:"bio,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

// Save inserts the User into the database.
func (u *User) Save(db *sql.DB) error {
	if db == nil {
		return fmt.Errorf("database connection is nil")
	}

	query := `
		INSERT INTO users (name, email, password, role, avatar_url, bio)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at;
	`

	row := db.QueryRow(
		query,
		u.Name,
		u.Email,
		u.Password,
		u.Role,
		u.AvatarURL,
		u.Bio,
	)

	err := row.Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		// Check for PostgreSQL unique violation
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" && pqErr.Constraint == "users_email_key" {
				return fmt.Errorf("email '%s' is already registered", u.Email)
			}
		}
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}
