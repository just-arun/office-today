package routes

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/just-arun/office-today/internals/pkg/fileupload"
)

// Fileupload for uploading file
func Fileupload(r *mux.Router) {
	fmt.Println("File upload route registered...")
	s := r.PathPrefix("/file-upload").Subrouter()

	s.HandleFunc(
		"",
		fileupload.UploadFile,
	).
		Methods("POST")
}
