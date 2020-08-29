package routes

import (
	"fmt"

	"github.com/gorilla/mux"
)

// Users register the routes for users
func Users(r *mux.Router) {
	fmt.Println("User route registered...")
	// s := r.PathPrefix("/user").Subrouter()
}
