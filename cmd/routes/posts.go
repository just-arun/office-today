package routes

import (
	"fmt"

	"github.com/gorilla/mux"
)

// Posts register the routes for posts
func Posts(r *mux.Router) {
	fmt.Println("Post route registered...")
	// s := r.PathPrefix("/posts").Subrouter()
}
