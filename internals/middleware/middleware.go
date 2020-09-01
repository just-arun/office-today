package middleware

import (
	"fmt"
	"io"
	"net/http"

	mContext "github.com/gorilla/context"
	"github.com/just-arun/office-today/internals/pkg/users/usertype"
)

// Auth authentication of user check if the users are logedin
func Auth(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Auth middleware...")
		next(w, r)
	}
}

// Owner authentication of user check if the users are logedin
func Owner(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Auth middleware...")
		next(w, r)
	}
}

// UserType authentication of user check if the users are logedin
func UserType(next func(http.ResponseWriter, *http.Request), userType []usertype.UserType) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("user type middleware...")
		usType := mContext.Get(r, "type")
		for _, uType := range userType {
			if usType == uType {
				next(w, r)
				return
			}
		}
		io.WriteString(w, "not ment")
		return
	}
}
