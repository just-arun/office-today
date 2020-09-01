package users

import (
	"net/http"
)

// GetUser get users
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("some shit"))
}
