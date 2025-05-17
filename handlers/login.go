package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Hritikpandey-ops/blog-site/repository"
	"github.com/Hritikpandey-ops/blog-site/utils"
)

// Login handles user login and returns an auth token with user details.
func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		// Decode request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Authenticate user
		user, err := repository.CheckLogin(db, req.Email, req.Password)
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		// Generate auth token (with user ID)
		token, err := utils.GenerateAuthToken(user.ID)
		if err != nil {
			http.Error(w, "Could not generate auth token", http.StatusInternalServerError)
			return
		}

		// Return success response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"user_id": user.ID.String(),
			"name":    user.Name,
			"email":   user.Email,
			"role":    user.Role,
			"token":   token,
			"message": "Login successful",
		})
	}
}
