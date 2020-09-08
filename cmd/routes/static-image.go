package routes

import (
	"flag"
	"net/http"

	"github.com/gorilla/mux"
)

// StaticFile for accessing static files
func StaticFile(r *mux.Router) {
	var dir string

	flag.StringVar(&dir, "dir", "./images/", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	// This will serve files under http://localhost:8000/static/<filename>
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
}
