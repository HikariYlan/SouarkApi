package main

import (
	"log"
	"net/http"
	"os"
	"souark/api/player"
	"souark/api/services/send"
	"souark/api/team"

	"github.com/joho/godotenv"
)

func apiVersion(res http.ResponseWriter, _ *http.Request) {
	err := godotenv.Load()
	if err != nil {
		panic("No environment file found")
	}

	version := os.Getenv("API_VERSION")

	send.Json(version, res, http.StatusOK)
}

func main() {
	// Router creation
	router := http.NewServeMux()

	// API version
	router.HandleFunc("/api/version", apiVersion)

	// Player router
	router.Handle("/api/player/", http.StripPrefix("/api/player", player.Router()))

	// Team router
	router.Handle("/api/team/", http.StripPrefix("/api/team", team.Router()))

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
