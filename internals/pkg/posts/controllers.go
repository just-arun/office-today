package posts

import (
	"encoding/json"
	"net/http"

	// gCtx "github.com/gorilla/context"
	"github.com/just-arun/office-today/internals/middleware/response"
)

// CreatePost for creating post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post Posts
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		response.Error(w, http.StatusBadGateway, err.Error())
		return
	}
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
	// postID := mux.Vars(r)["id"]
	return
}

// GetAllPost for getting all post
func GetAllPost(w http.ResponseWriter, r *http.Request) {
	// page := r.URL.Query()["page"]
	return
}

// GetUserPost for getting all user post
func GetUserPost(w http.ResponseWriter, r *http.Request) {
	// page := r.URL.Query()["page"]
	// userID := mux.Vars(r)["id"]
	return
}

//DisablePost for disable user post
func DisablePost(w http.ResponseWriter, r *http.Request) {
	// postID := mux.Vars(r)["id"]
	return
}
