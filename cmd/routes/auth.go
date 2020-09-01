package routes

import (
	"fmt"

	"github.com/just-arun/office-today/internals/middleware"

	"github.com/just-arun/office-today/internals/pkg/auth"

	"github.com/gorilla/mux"
)

// Auth register the routes for auth
func Auth(r *mux.Router) {
	fmt.Println("Auth route register...")
	s := r.PathPrefix("/auth").Subrouter()

	// register password route
	s.HandleFunc("/register",
		auth.Register,
	).Methods("POST")

	// login users
	s.HandleFunc("/login",
		auth.Login,
	).Methods("POST")

	// forgot password for users
	s.HandleFunc("/forgot-password",
		auth.ForgotPassword,
	).Methods("POST")

	// reset password for users
	s.HandleFunc("/reset-password",
		auth.ResetPassword,
	).Methods("POST")

	// reset refresh token for users
	s.HandleFunc("/refresh-token",
		auth.RefreshToken,
	).Methods("PATCH")

	// update password for users
	s.HandleFunc("/update-password",
		middleware.Auth(
			auth.UpdatePassword,
		),
	).Methods("PATCH")
}
