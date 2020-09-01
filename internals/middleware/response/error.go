package response

import (
	"encoding/json"
	"net/http"
)

type errorData struct {
	Error map[string]interface{} `json:"error"`
}

// Error for failure response
func Error(w http.ResponseWriter, status int, errStr string) {
	errData := &errorData{
		Error: map[string]interface{}{
			"message": errStr,
		},
	}
	data, _ := json.Marshal(errData)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}
