package middleware

import (
	"fmt"
	"net/http"

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
		fmt.Println("Auth middleware...")
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

		gCtx.Set(r, "uid", user.ID)
		gCtx.Set(r, "email", user.Email)
		gCtx.Set(r, "type", user.Type)

		next(w, r)
		return
	}
}

// Owner authentication of user check if the users are logedin
func Owner(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Auth middleware...")
		next(w, r)
	}
}

// UserType authentication of user check if the users are logedin
func UserType(next func(http.ResponseWriter, *http.Request), userType ...usertype.UserType) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("user type middleware...")
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
