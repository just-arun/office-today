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

	s.HandleFunc("/{id}/removelike",
		middleware.Auth(
			posts.RemoveLike,
		),
	).Methods("POST")

	s.HandleFunc("/{id}/comments",
		middleware.Auth(
			posts.GetComments,
		),
	).Methods("GET")

	s.HandleFunc("/{id}/enquiry",
		middleware.Auth(
			posts.CreateEnquiry,
		),
	).Methods("POST")

	s.HandleFunc("/{id}/enquiry",
		middleware.Auth(
			posts.GetEnquiry,
		),
	).Methods("GET")

	s.HandleFunc("/tag/get",
		middleware.Auth(
			posts.GetTags,
		),
	).Methods("GET")

	s.HandleFunc("/tag/create",
		middleware.Auth(
			middleware.UserType(
				posts.CreateTags,
				usertype.Admin,
			),
		),
	).Methods("POST")

	s.HandleFunc("/tag/delete/{id}",
		middleware.Auth(
			middleware.UserType(
				posts.DeleteTag,
				usertype.Admin,
			),
		),
	).Methods("DELETE")

	s.HandleFunc("/search/post", 
		middleware.Auth(
			posts.SearchPost,
		),
	).Methods("GET")

}
