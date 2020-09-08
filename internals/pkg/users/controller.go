package users

import (
	"net/http"
	"strconv"

	"github.com/just-arun/office-today/internals/middleware/response"
	"github.com/just-arun/office-today/internals/pkg/comments"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/gorilla/mux"
)

// GetUsers get users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	page := mux.Vars(r)["id"]

	count := 0

	if len(page) > 0 {
		num, err := strconv.Atoi(page)
		if err != nil {
			response.Error(w, http.StatusBadGateway, err.Error())
			return
		}
		count = num
	}

	user, err := GetAll(
		bson.M{},
		count,
	)

	if err != nil {
		response.Error(w, http.StatusBadGateway, err.Error())
		return
	}

	response.Success(w, r,
		http.StatusOK,
		user,
	)
	return
}

// GetComments for getting user comments
func GetComments(w http.ResponseWriter, r *http.Request) {
	uID := mux.Vars(r)["id"]
	var comment []*comments.Comments
	err := GetUserComments(uID, comment)
	if err != nil {
		response.Error(w, http.StatusBadGateway, err.Error())
		return
	}
	response.Success(
		w, r,
		http.StatusOK,
		comment,
	)
	return
}