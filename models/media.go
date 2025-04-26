package models

import "time"

type Media struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	FileName  string    `json:"file_name"`
	FileURL   string    `json:"file_url"`
	MimeType  *string   `json:"mime_type,omitempty"` // Nullable
	Size      *int64    `json:"size,omitempty"`      // Nullable
	CreatedAt time.Time `json:"created_at"`
}
