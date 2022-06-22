package helpers

import (
	"net/http"
)

// RespondJSON translates an interface to json for response
func RespondError(w http.ResponseWriter, err error) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(err.Error()))
}
