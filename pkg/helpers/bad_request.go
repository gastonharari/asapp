package helpers

import (
	"net/http"
)

// RespondJSON translates an interface to json for response
func RespondBadRequest(w http.ResponseWriter, message string) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(message))
}
