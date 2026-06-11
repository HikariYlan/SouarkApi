package player

import (
	"net/http"
)

// Returns a configured router with prefix /api/player associated routes (without the prefix)
func Router() *http.ServeMux {
	return controller.NewRouter()
}
