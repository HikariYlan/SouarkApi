package game

import (
	"net/http"
)

func Router() *http.ServeMux {
	router := controller.NewRouter()

	router.HandleFunc("POST /new", NewGame)

	return router
}
