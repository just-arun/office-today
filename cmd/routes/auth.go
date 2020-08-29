package routes

import (
	"fmt"

	"github.com/gorilla/mux"
)

// Auth register the routes for auth
func Auth(r *mux.Router) {
	fmt.Println("Auth route register...")
	// s := r.PathPrefix("/auth").Subrouter()
}
