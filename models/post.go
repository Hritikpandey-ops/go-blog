package models

import (
	"time"

	"github.com/Hritikpandey-ops/blog-site/enums"
)

type Post struct {
	ID            int              `json:"id"`
	UserID        int              `json:"user_id"`
	Title         string           `json:"title"`
	Slug          string           `json:"slug"`
	Excerpt       *string          `json:"excerpt,omitempty"`
	Content       string           `json:"content"`
	FeaturedImage *string          `json:"featured_image,omitempty"`
	Status        enums.PostStatus `json:"status"`
	PublishedAt   *time.Time       `json:"published_at,omitempty"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`
}
