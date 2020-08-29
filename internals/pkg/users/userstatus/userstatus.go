package userstatus

// UserStatus users status
type UserStatus int

const (
	// Disabled user status
	Disabled UserStatus = iota
	// Active user status
	Active
	// ForgotPssword user status
	ForgotPssword
	// ResetPassword user status
	ResetPassword
)
