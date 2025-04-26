package enums

type UserRole string

const (
	RoleReader UserRole = "reader"
	RoleAuthor UserRole = "author"
	RoleAdmin  UserRole = "admin"
)
