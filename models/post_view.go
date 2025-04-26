package models

import "time"

type PostView struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	UserID    *int      `json:"user_id,omitempty"`    // Nullable
	IPAddress *string   `json:"ip_address,omitempty"` // Nullable
	CreatedAt time.Time `json:"created_at"`
}
