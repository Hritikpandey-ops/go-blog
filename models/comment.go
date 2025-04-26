package models

import "time"

type Comment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	UserID    *int      `json:"user_id,omitempty"` // Nullable
	Name      *string   `json:"name,omitempty"`    // Nullable
	Email     *string   `json:"email,omitempty"`   // Nullable
	Content   string    `json:"content"`
	ParentID  *int      `json:"parent_id,omitempty"` // Nullable
	CreatedAt time.Time `json:"created_at"`
}
