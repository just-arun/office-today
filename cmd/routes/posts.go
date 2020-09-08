package routes

import (
	"fmt"

	"github.com/just-arun/office-today/internals/middleware"
	"github.com/just-arun/office-today/internals/middleware/ownerarea"
	"github.com/just-arun/office-today/internals/pkg/posts"
	"github.com/just-arun/office-today/internals/pkg/users/usertype"

	"github.com/gorilla/mux"
)

// Posts register the routes for posts
func Posts(r *mux.Router) {
	fmt.Println("Post route registered...")
	s := r.PathPrefix("/posts").Subrouter()

	// create post
	s.HandleFunc("",
		middleware.Auth(
			middleware.UserType(
				posts.CreatePost,
				usertype.Admin,
				usertype.Publisher,
			),
		),
	).
		Methods("POST")

	s.HandleFunc("",
		middleware.Auth(
			posts.GetAllPost,
		),
	).
		Methods("GET")

	s.HandleFunc("/{id}",
		middleware.Auth(
			middleware.Owner(
				posts.DisablePost,
				ownerarea.Post,
			),
		),
	).
		Methods("DELETE")

	s.HandleFunc("/{id}",
		middleware.Auth(
			middleware.Owner(
				posts.UpdateOne,
				ownerarea.Post,
			),
		),
	).
		Methods("PUT")

	s.HandleFunc("/{id}",
		middleware.Auth(
			posts.GetOnePost,
		),
	).
		Methods("GET")

	s.HandleFunc("/{id}/addlike",
		middleware.Auth(
			posts.AddLike,
		),
	).Methods("POST")

}
