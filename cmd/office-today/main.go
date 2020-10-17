package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/just-arun/office-today/internals/boot/database"

	"github.com/just-arun/office-today/internals/boot/config"

	"github.com/gorilla/mux"
	"github.com/just-arun/office-today/cmd/routes"
)

func main() {
	fmt.Println("Setting up application...")
	config.Init()
	// Init database
	database.Init()
	// defining multiplexer
	r := mux.NewRouter()
	// Regestering routes
	routes.Auth(r)
	routes.Users(r)
	routes.Posts(r)
	routes.Comments(r)
	routes.Fileupload(r)
	routes.StaticFile(r)

	port := fmt.Sprintf("%v%v", config.AppHost, config.Port)

	fmt.Printf("server started at port http://%v\n", port)

	fmt.Println(port)

	fmt.Println(time.Now())
	// srv := &http.Server{
	// 	Handler:      r,
	// 	Addr:         port,
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }

	// log.Fatal(srv.ListenAndServe(), "server terminated")
	log.Fatal(http.ListenAndServe(port, r), "Server crashed")
}
