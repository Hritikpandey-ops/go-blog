package enums

type PostStatus string

const (
	StatusDraft     PostStatus = "draft"
	StatusPublished PostStatus = "published"
	StatusArchived  PostStatus = "archived"
)
