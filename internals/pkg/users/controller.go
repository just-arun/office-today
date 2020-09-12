package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	gCtx "github.com/gorilla/context"
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

// CreateUser for creating user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user Users
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.Error(w, http.StatusBadGateway, err.Error())
		return
	}

	uID, err := CreateUserService(user)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(
		w, r,
		http.StatusOK,
		map[string]interface{}{
			"id": uID,
		},
	)
	return
}

// UpdateUser for updating user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var userID = mux.Vars(r)["id"]

	var user UpdateUserStruct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.Error(w, http.StatusBadGateway, err.Error())
		return
	}

	fmt.Println(user)

	err := UpdateUserService(userID, bson.M{
		"$set": user,
	})
	if err != nil {
		response.Error(w, http.StatusBadGateway, err.Error())
		return
	}

	response.Success(
		w, r,
		http.StatusOK,
		map[string]interface{}{
			"ok": 1,
		},
	)
	return
}

// BookmarkHandle for adding to bookmark
func BookmarkHandle(w http.ResponseWriter, r *http.Request) {
	uID := gCtx.Get(r, "uid").(string)
	bookmarkType := r.URL.Query()["type"][0]

	var bookmark Bookmark
	if err := json.NewDecoder(r.Body).Decode(&bookmark); err != nil {
		response.Error(w, http.StatusBadGateway, err.Error())
		return
	}
	var err error
	if bookmarkType == "add" {
		err = AddBookmarkService(uID, bookmark.ID)
	} else {
		err = RemoveBookmarkService(uID, bookmark.ID)
	}

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(
		w, r,
		http.StatusOK,
		map[string]interface{}{
			"ok": 1,
		},
	)
	return
}
