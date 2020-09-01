package auth

import (
	"errors"

	"github.com/just-arun/office-today/internals/pkg/users/userstatus"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/just-arun/office-today/internals/util/password"

	"github.com/just-arun/office-today/internals/pkg/users"
	"github.com/just-arun/office-today/internals/pkg/users/usertype"
)

// RegisterService user
func RegisterService(u *users.Users) (interface{}, error) {
	dbUser, _ := users.GetOne(bson.M{
		"email": u.Email,
	})
	if dbUser != nil {
		return nil, errors.New("User already exist")
	}

	u.Type = usertype.Audience
	u.Status = userstatus.Active
	userID, err := u.Save()
	if err != nil {
		return nil, err
	}
	return userID, nil
}

// LoginService for user
func LoginService(login *LoginDto) (map[string]interface{}, error) {
	// get User from database
	user, err := users.GetOne(bson.M{
		"email": login.Email,
		"status": bson.M{
			"$ne": userstatus.Disabled,
		}})

	if err != nil {
		return nil, errors.New("invalided credentials")
	}

	if !password.Compare(login.Password, user.Password) {
		return nil, errors.New("invalided credentials")
	}

	return map[string]interface{}{
		"email":    user.Email,
		"userName": user.UserName,
		"_id":      user.ID,
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

// RefreshTokenService return refresh token
func RefreshTokenService(token *RefreshTokenDto) (map[string]interface{}, error) {
	user, err := users.GetOne(bson.M{
		"refresh_token": token.RefreshToken,
	})

	if err != nil {
		return nil, errors.New("invalided token")
	}

	return map[string]interface{}{
		"uid": user.ID.Hex(),
	}, nil
}
