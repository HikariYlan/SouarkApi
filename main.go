package main

import (
	"log"
	"net/http"
	"souark/api/player"
)

func main() {
	// Router creation
	router := http.NewServeMux()

	// Player router
	router.Handle("/api/player/", http.StripPrefix("/api/player", player.Router()))

	// OpenAPI documentation
	fs := http.FileServer(http.Dir("./api_doc"))
	router.Handle("/api_doc/", http.StripPrefix("/api_doc", fs))

	// Server configuration
	server := http.Server{
		Addr:    ":7777",
		Handler: router,
	}

	log.Println("Local server started : http://localhost:7777")

	// Server starting
	err := server.ListenAndServe()

	log.Fatalln(err)
}
