package routes

import (
	"fmt"

	"github.com/just-arun/office-today/internals/middleware"

	"github.com/just-arun/office-today/internals/pkg/comments"

	"github.com/gorilla/mux"
)

// Comments register the routes for comment
func Comments(r *mux.Router) {
	fmt.Println("User route registered...")
	s := r.PathPrefix("/comments").Subrouter()

	// Create comment
	s.HandleFunc("",
		middleware.Auth(
			comments.CreateComment,
		),
	).
		Methods("POST")
}
