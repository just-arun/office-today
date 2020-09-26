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
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"
)

// GetUsers get users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query()["page"]
	userID := gCtx.Get(r, "uid").(string)
	uID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		response.Error(w, http.StatusBadGateway, err.Error())
		return
	}

	count := 0

	if len(page) > 0 {
		num, err := strconv.Atoi(page[0])
		if err != nil {
			response.Error(w, http.StatusBadGateway, err.Error())
			return
		}
		count = num
	}

	user, err := GetAll(
		bson.M{
			"_id": bson.M{
				"$ne": uID,
			},
		},
		count,
	)

	if err != nil {
		response.Error(w, http.StatusBadGateway, err.Error())
		return
	}

	if user == nil {
		user = []*UsersStruct{}
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

// GetUserProfile from getting user profile
func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	uID := gCtx.Get(r, "uid").(string)
	user, err := GetUserProfileService(uID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(user)
	if user == nil {
		user = &UsersStruct{}
	}

	response.Success(
		w, r,
		http.StatusOK,
		user,
	)
	return
}

// SearchUsers for searching user
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query()["key"]
	var query = ""
	if len(key) > 0 {
		if key[0] != "" {
			query = key[0]
		}
	}
	fmt.Printf("query %v\n", query)
	reuslt, err := SearchUserService(query)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	if reuslt == nil {
		reuslt = []SearchStruct{}
	}

	response.Success(
		w, r,
		http.StatusOK,
		reuslt,
	)
	return
}

// GetUserPosts for getting users post
func GetUserPosts(w http.ResponseWriter, r *http.Request) {
	uID := mux.Vars(r)["id"]
	page := r.URL.Query()["page"]

	count := 0
	if len(page) > 0 {
		if page[0] != "" {
			num, err := strconv.Atoi(page[0])
			if err != nil {
				response.Error(w, http.StatusUnprocessableEntity, err.Error())
				return
			}
			count = num
		}
	}

	result, err := GetUserPostServices(count, uID)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(
		w, r,
		http.StatusOK,
		result,
	)
	return
}

// GetOneUser for getting one user
func GetOneUser(w http.ResponseWriter, r *http.Request) {
	user := mux.Vars(r)["id"]
	uID, err := primitive.ObjectIDFromHex(user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	result, err := GetOneUserService(bson.M{
		"_id": uID,
	})

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(
		w, r,
		http.StatusOK,
		result,
	)
	return
}


