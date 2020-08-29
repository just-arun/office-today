package routes

import (
	"fmt"

	"github.com/gorilla/mux"
)

// Comments register the routes for comment
func Comments(r *mux.Router) {
	fmt.Println("User route registered...")
	// s := r.PathPrefix("/comments").Subrouter()
}
