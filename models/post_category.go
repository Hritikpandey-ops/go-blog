package models

type PostCategory struct {
	ID         int `json:"id"`
	PostID     int `json:"post_id"`
	CategoryID int `json:"category_id"`
}
