package users

import (
	"net/http"
	"strconv"

	"github.com/just-arun/office-today/internals/middleware/response"
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
