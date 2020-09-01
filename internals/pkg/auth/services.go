package auth

import (
	"github.com/just-arun/office-today/internals/util/password"
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
func LoginService(login *LoginDto) (map[string]interface{}, error) {
	// get User from database
	user, err := users.GetOne(map[string]interface{}{
		"email": login.Email,
	})

	if err != nil {
		return nil, err
	}

	if !password.Compare(login.Password, user.Password) {
		return nil, err
	}

	return map[string]interface{}{
		"email": user.Email,
		"userName": user.UserName,
		"_id": user.ID,
	}, nil
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
