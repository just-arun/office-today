package auth

import (
	"io"
	"net/http"
)

// Register user
func Register(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "bullshit")
}

// Login for user
func Login(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is login shit")
}

// ForgotPassword password status set
func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "forgot password shit")
}

// ResetPassword reset password
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "reset password")
}

// UpdatePassword update password
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "update password")
}
