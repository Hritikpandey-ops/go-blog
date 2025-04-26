package models

type PostTag struct {
	ID     int `json:"id"`
	PostID int `json:"post_id"`
	TagID  int `json:"tag_id"`
}
