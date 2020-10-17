package auth

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/just-arun/office-today/internals/boot/collections"
	"github.com/just-arun/office-today/internals/boot/config"
	"github.com/just-arun/office-today/internals/util/aesencryption"
	"github.com/just-arun/office-today/internals/util/message"

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
		},
	})
	fmt.Println(user, login.Email)
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
func ForgotPasswordService(email string) error {

	otp := rand.Intn(999999)

	_, err := collections.
		User().
		UpdateOne(context.TODO(),
			bson.M{"email": email},
			bson.M{
				"$set": bson.M{
					"otp": otp,
				},
			})
	if err != nil {
		if err != nil {
			fmt.Println("error", err.Error())
			return err
		}
		return err
	}

	msg := "OTP: " + fmt.Sprint(otp) + " for loggin to offer today app "

	err = message.Mail(email, "Forgot Password", msg)
	fmt.Println(msg)
	if err != nil {
		fmt.Println("error", err.Error())
		return err
	}
	return nil
}

// ResetPasswordService reset password
func ResetPasswordService() {

}

// UpdatePasswordService update password
func UpdatePasswordService(email string, otp int) error {
	var user users.Users
	err := collections.User().FindOne(context.TODO(), bson.M{"email": email, "otp": otp}).
	Decode(&user)
	if err != nil {
		return err
	}

	pwd, err := password.Encrypt(user.Password)
	if err != nil {
		return err
	}

	_, err = collections.
		User().
		UpdateOne(context.TODO(),
			bson.M{"email": email},
			bson.M{
				"$set": bson.M{
					"otp": nil,
					"password": pwd,
				},
			})
	if err != nil {
		if err != nil {
			fmt.Println("error", err.Error())
			return err
		}
		return err
	}
	return nil
}

// RefreshTokenService return refresh token
func RefreshTokenService(token *RefreshTokenDto) (map[string]interface{}, error) {

	tokenID := aesencryption.Decrypt([]byte(config.AESSecret), token.RefreshToken)

	uID, err := primitive.ObjectIDFromHex(tokenID)

	if err != nil {
		return nil, errors.New("invalided token")
	}

	user, err := users.GetOne(bson.M{
		"_id":           uID,
		"refresh_token": token.RefreshToken,
	})

	if err != nil {
		return nil, errors.New("invalided token")
	}

	return map[string]interface{}{
		"uid": user.ID.Hex(),
	}, nil
}
