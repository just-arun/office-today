package posts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/just-arun/office-today/internals/pkg/posts/poststatus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	gCtx "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/just-arun/office-today/internals/middleware/response"
)

// CreatePost for creating post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var createPost CreatePostDto
	if err := json.NewDecoder(r.Body).Decode(&createPost); err != nil {
		response.Error(w, http.StatusBadGateway, err.Error())
		return
	}

	var post Posts
	post.Title = createPost.Title
	post.Description = createPost.Description
	post.ImageURL = createPost.ImageURL

	userID := gCtx.Get(r, "uid")
	postID, err := post.Save(userID.(string))
	if err != nil {
		response.Error(w, http.StatusBadGateway, err.Error())
		return
	}
	response.Success(w, r,
		http.StatusCreated,
		map[string]interface{}{
			"id": postID,
		},
	)
	return
}

// UpdateOne for updating post
func UpdateOne(w http.ResponseWriter, r *http.Request) {
	// postID := mux.Vars(r)["id"]
	var post Posts
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		response.Error(w, http.StatusBadGateway, err.Error())
		return
	}
}

// GetOnePost for getting one post
func GetOnePost(w http.ResponseWriter, r *http.Request) {
	postID := mux.Vars(r)["id"]
	pID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	res, err := GetOne(bson.M{
		"_id": pID,
	})
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	response.Success(w, r,
		http.StatusOK,
		res,
	)
	return
}

// GetAllPost for getting all post
func GetAllPost(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query()["page"]
	userID := r.URL.Query()["user"]

	filter := bson.M{}
	filter["status"] = bson.M{
		"$ne": poststatus.Deleted,
	}

	if len(userID) > 0 {
		uID, err := primitive.ObjectIDFromHex(userID[0])
		if err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err.Error())
			return
		}
		filter["user_id"] = uID
	}

	var count int

	if len(page) > 0 {
		fmt.Println("STUFF", len(page))
		num, err := strconv.Atoi(page[0])
		if err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err.Error())
			return
		}
		count = num
	} else {
		count = 0
	}

	posts, err := GetAll(filter, count)

	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	response.Success(w, r,
		http.StatusOK,
		posts,
	)
	return
}

// GetMyPost for getting all my post
func GetMyPost(w http.ResponseWriter, r *http.Request) {
	// page := r.URL.Query()["page"]
	// userID := gCtx.Get(r)["uid"]
	return
}

//DisablePost for disable user post
func DisablePost(w http.ResponseWriter, r *http.Request) {
	// postID := mux.Vars(r)["id"]
	return
}
