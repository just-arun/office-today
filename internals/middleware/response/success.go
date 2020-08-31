package response

import (
	"encoding/json"
	"net/http"

	"github.com/just-arun/office-today/internals/pkg/users"

	"github.com/just-arun/office-today/internals/util/tokens"

	gCtx "github.com/gorilla/context"
)

type successData struct {
	Status       int                    `json:"status"`
	Data         map[string]interface{} `json:"data"`
	AccessToken  string                 `json:"accessToken"`
	RefreshToken string                 `json:"refreshToken"`
	ResetPwd     bool                   `json:"resetPwd"`
}

// Success response
func Success(
	w http.ResponseWriter,
	r *http.Request,
	status int,
	data map[string]interface{},
) {
	var resData successData
	resData.Status = status
	resData.Data = data

	refresh := gCtx.Get(r, "refresh")
	if refresh == true {
		id := gCtx.Get(r, "uid")
		access, err := tokens.GenerateToken(id.(string), tokens.AccessToken)
		if err != nil {
			Error(w, http.StatusInternalServerError, err.Error())
			return
		}
		refresh, err := tokens.GenerateToken(id.(string), tokens.RefreshToken)
		if err != nil {
			Error(w, http.StatusInternalServerError, err.Error())
			return
		}
		err = users.UpdateRefreshToken(id.(string), refresh)
		if err != nil {
			Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		resData.AccessToken = access
		resData.RefreshToken = refresh

	}

	jsonData, err := json.Marshal(resData)

	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)
}
