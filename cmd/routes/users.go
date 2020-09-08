package routes

import (
	"fmt"

	"github.com/just-arun/office-today/internals/pkg/users/usertype"

	"github.com/just-arun/office-today/internals/middleware"
	"github.com/just-arun/office-today/internals/pkg/users"

	"github.com/gorilla/mux"
)

// Users register the routes for users
func Users(r *mux.Router) {
	fmt.Println("User route registered...")
	s := r.PathPrefix("/user").Subrouter()

	// GetUsers for getting all users
	s.HandleFunc("",
		middleware.Auth(
			users.GetUsers,
		),
	).Methods("GET")

	// UpdateUser for updating user
	s.HandleFunc("/{id}",
		middleware.Auth(
			users.UpdateUser,
		),
	).Methods("PUT")

	s.HandleFunc("",
		middleware.Auth(
			middleware.UserType(
				users.CreateUser,
				usertype.Admin,
			),
		),
	).
    Methods("POST")
    
  s.HandleFunc("/{id}/posts/add",
    middleware.Auth(
      middleware.Owner(
        users.AddBookmark,
      )
    )
  ).Methods("POST")

  s.HandleFunc("/{id}/posts/remove",
    middleware.Auth(
      middleware.Owner(
        users.RemoveBookmark,
      )
    )
  ).Methods("POST")
}
