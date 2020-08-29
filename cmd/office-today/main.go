package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/just-arun/office-today/internals/boot/config"

	"github.com/gorilla/mux"
	"github.com/just-arun/office-today/cmd/routes"
)

func main() {
	fmt.Println("server started at port", config.Port)
	r := mux.NewRouter()
	// Regestering routes
	routes.Auth(r)
	routes.Users(r)
	routes.Posts(r)
	routes.Comments(r)

	log.Fatal(http.ListenAndServe(config.Port, r), "server terminated")
}
