package auth

// LoginDto for login requrest
type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
