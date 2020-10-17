package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/just-arun/office-today/internals/middleware/response"

	"github.com/just-arun/office-today/internals/pkg/users"

	gCtx "github.com/gorilla/context"
)

// Register user
func Register(w http.ResponseWriter, r *http.Request) {
	var user users.Users
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.Error(w, http.StatusBadGateway, err.Error())
		return
	}
	userID, err := RegisterService(&user)

	if err != nil {
		response.Error(w, http.StatusBadGateway, err.Error())
		return
	}

	gCtx.Set(r, "refresh", true)
	gCtx.Set(r, "uid", userID)

	response.Success(
		w, r,
		http.StatusCreated,
		map[string]interface{}{
			"_id":      userID,
			"email":    user.Email,
			"userName": user.UserName,
		})

}

// Login for user
func Login(w http.ResponseWriter, r *http.Request) {
	var login LoginDto
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		response.Error(w, http.StatusForbidden, err.Error())
		return
	}
	fmt.Println(login)
	result, err := LoginService(&login)
	if err != nil {
		response.Error(w, http.StatusForbidden, err.Error())
		return
	}
	fmt.Println("[ID]", result["_id"], err)

	gCtx.Set(r, "refresh", true)
	gCtx.Set(r, "uid", result["_id"].(primitive.ObjectID).Hex())

	response.Success(
		w, r,
		http.StatusOK,
		result,
	)
}

// ForgotPassword password status set
func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var forGotPwd ForgotPasswordDto
	if err := json.NewDecoder(r.Body).Decode(&forGotPwd); err != nil {
		response.Error(w, http.StatusBadGateway, err.Error())
		return
	}
	err := ForgotPasswordService(forGotPwd.Email)
	if err != nil {
		response.Error(w, http.StatusBadGateway, err.Error())
		return
	}
	response.Success(w, r,
		http.StatusOK,
		map[string]interface{}{
			"ok": 1,
		},
	)
	return
}

// ResetPassword reset password
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var user users.Users
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		io.WriteString(w, "update password")
		return
	}
	users.ResetPassword(user.Email)
	io.WriteString(w, "update password")
}

// UpdatePassword update password
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "update password")
}

// RefreshToken update access token
func RefreshToken(w http.ResponseWriter, r *http.Request) {
	var token RefreshTokenDto
	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		response.Error(w, http.StatusForbidden, err.Error())
		return
	}
	fmt.Println("[token]", token)
	tokenData, err := RefreshTokenService(&token)
	if err != nil {
		response.Error(w, http.StatusForbidden, err.Error())
		return
	}

	gCtx.Set(r, "refresh", true)
	gCtx.Set(r, "uid", tokenData["uid"])

	response.Success(
		w, r,
		http.StatusOK,
		map[string]interface{}{
			"ok": 1,
		})
}
