package repository

import (
	"database/sql"
	"errors"

	"github.com/Hritikpandey-ops/blog-site/models"
	"golang.org/x/crypto/bcrypt"
)

func CheckLogin(db *sql.DB, email, password string) (*models.User, error) {
	user, err := GetUserByEmail(db, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

func GetUserByEmail(db *sql.DB, email string) (*models.User, error) {
	query := `
		SELECT id, name, email, password, role, avatar_url, bio, created_at, updated_at
		FROM users
		WHERE email = $1
	`
	var user models.User
	err := db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.AvatarURL,
		&user.Bio,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
