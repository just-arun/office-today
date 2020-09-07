package comments

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	mCtx "github.com/gorilla/context"
	"github.com/just-arun/office-today/internals/middleware/response"
)

// CreateComment for creating comment
func CreateComment(w http.ResponseWriter, r *http.Request) {
	userID := mCtx.Get(r, "uid")
	var comment Comments
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		response.Error(w, http.StatusBadGateway, err.Error())
		return
	}
	uID, err := comment.Save(userID.(string))
	if err != nil {
		response.Error(w, http.StatusBadGateway, err.Error())
		return
	}
	response.Success(
		w, r,
		http.StatusCreated,
		map[string]interface{}{
			"id": uID,
		},
	)
}

// DeleteComment for deleting comment
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	commentID := mux.Vars(r)["id"]
	if err := DeleteCommentService(commentID); err != nil {
		response.Error(w, http.StatusBadGateway, err.Error())
		return
	}
	response.Success(
		w, r,
		http.StatusCreated,
		map[string]interface{}{
			"id": commentID,
		},
	)
	return
}
