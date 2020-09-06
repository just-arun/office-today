package auth

// LoginDto for login requrest
type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}


// RefreshTokenDto get refresh token
type RefreshTokenDto struct {
	RefreshToken string `json:"refreshToken"`
}

// ForgotPasswordDto for getting email
type ForgotPasswordDto struct {
	Email string `json:"email"`
}