package middleware

import (
	"fmt"
	"net/http"

	"github.com/just-arun/office-today/internals/middleware/ownerarea"

	"github.com/just-arun/office-today/internals/pkg/users/userstatus"

	"github.com/just-arun/office-today/internals/pkg/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/just-arun/office-today/internals/util/stringutil"

	"github.com/just-arun/office-today/internals/util/tokens"

	gCtx "github.com/gorilla/context"
	"github.com/just-arun/office-today/internals/pkg/users/usertype"
)

// Auth authentication of user check if the users are logedin
func Auth(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := tokens.GetTokenFromHeader(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		_, claim, err := tokens.DecodeJWTToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		uID, err := stringutil.StringFromHash(claim["id"].(string))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userUID, err := primitive.ObjectIDFromHex(uID)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		user, err := users.GetOne(bson.M{
			"_id": userUID,
			"status": bson.M{
				"$ne": userstatus.Disabled,
			},
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		gCtx.Set(r, "uid", user.ID.Hex())
		gCtx.Set(r, "email", user.Email)
		gCtx.Set(r, "type", user.Type)

		next(w, r)
		return
	}
}

// Owner authentication of user check if the users are logedin
func Owner(next func(http.ResponseWriter, *http.Request), ownerAccess ownerarea.OwnerArea) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// accessID := mux.Vars(r)["id"]
		// userID := gCtx.Get(r, "uid")
		// userType := gCtx.Get(r, "type")

		// // check on user table
		// isOwner := accessID == userID

		switch ownerAccess {
		case ownerarea.User:
			fmt.Println("owner user")
			next(w, r)
			break
		case ownerarea.Post:
			fmt.Println("owner post")
			next(w, r)
			break
		case ownerarea.Like:
			fmt.Println("owner like")
			next(w, r)
			break
		case ownerarea.Comment:
			fmt.Println("owner comment")
			next(w, r)
			break
		case ownerarea.Bookmark:
			fmt.Println("owner bookmark")
			next(w, r)
			break
		}
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// UserType authentication of user check if the users are logedin
func UserType(next func(http.ResponseWriter, *http.Request), userType ...usertype.UserType) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		usType := gCtx.Get(r, "type")
		for _, uType := range userType {
			if usType == uType {
				next(w, r)
				return
			}
		}
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}
