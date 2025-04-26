package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Hritikpandey-ops/blog-site/enums"
	"github.com/Hritikpandey-ops/blog-site/models"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest

		// Parse JSON body
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Validate input
		req.Name = strings.TrimSpace(req.Name)
		req.Email = strings.TrimSpace(req.Email)
		req.Password = strings.TrimSpace(req.Password)

		if req.Name == "" || req.Email == "" || req.Password == "" {
			http.Error(w, "Name, Email and Password are required", http.StatusBadRequest)
			return
		}

		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}

		// Create User object
		user := &models.User{
			Name:     req.Name,
			Email:    req.Email,
			Password: string(hashedPassword),
			Role:     enums.RoleReader,
		}

		// Save to database
		if err := user.Save(db); err != nil {
			http.Error(w, "Failed to register user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond success
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "User registered successfully",
			"user_id": user.ID,
		})
	}
}
