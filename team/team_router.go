package team

import (
	"net/http"
)

func Router() *http.ServeMux {
	router := controller.NewRouter()

	router.HandleFunc("POST /new", NewTeam)

	return router
}
