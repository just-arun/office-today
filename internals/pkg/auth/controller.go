package auth

import (
	"encoding/json"
	"io"
	"net/http"

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
