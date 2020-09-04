package routes

import (
	"fmt"

	"github.com/just-arun/office-today/internals/middleware"
	"github.com/just-arun/office-today/internals/pkg/users"

	"github.com/gorilla/mux"
)

// Users register the routes for users
func Users(r *mux.Router) {
	fmt.Println("User route registered...")
	s := r.PathPrefix("/user").Subrouter()

	// GetUsers for getting all sers
	s.HandleFunc("",
		middleware.Auth(
			users.GetUsers,
		),
	).Methods("GET")
}
