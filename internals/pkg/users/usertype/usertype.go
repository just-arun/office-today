package usertype

// UserType user type
type UserType int

const (
	// Audience user type
	Audience UserType = iota
	// Publisher user type
	Publisher
	// Admin user type
	Admin
)