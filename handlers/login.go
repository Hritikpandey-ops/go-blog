package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Hritikpandey-ops/blog-site/repository"
	"github.com/Hritikpandey-ops/blog-site/utils"
)

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		user, err := repository.CheckLogin(db, req.Email, req.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := utils.GenerateAuthToken()
		if err != nil {
			http.Error(w, "Could not generate auth token", http.StatusInternalServerError)
			return
		}

		// Return response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"user_id": user.ID,
			"name":    user.Name,
			"email":   user.Email,
			"role":    user.Role,
			"token":   token,
			"message": "Login successful",
		})
	}
}
