package auth

import (
	"errors"

	"github.com/just-arun/office-today/internals/pkg/users"
	"github.com/just-arun/office-today/internals/pkg/users/usertype"
)

// RegisterService user
func RegisterService(u *users.Users) (interface{}, error) {
	dbUser, _ := users.GetOne(map[string]interface{}{
		"email": u.Email,
	})
	if dbUser != nil {
		return nil, errors.New("User already exist")
  }
  
	u.UserType = usertype.Audience
	userID, err := u.Save()
	if err != nil {
		return nil, err
	}
	return userID, nil
}

// LoginService for user
func LoginService() {

}

// ForgotPasswordService password status set
func ForgotPasswordService() {

}

// ResetPasswordService reset password
func ResetPasswordService() {

}

// UpdatePasswordService update password
func UpdatePasswordService() {

}
