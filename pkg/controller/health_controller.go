package controller

import (
	"net/http"

	"github.com/challenge/pkg/helpers"
	"github.com/challenge/pkg/models"
	"github.com/challenge/pkg/repository"
)

// Check returns the health of the service and DB
func (h Handler) Check(w http.ResponseWriter, r *http.Request) {
	// TODO: Check service health. Feel free to add any check you consider necessary
	db := repository.InitDb()
	if db == nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("error conecting to db"))
		return
	}
	repository.InitRepo(db)
	helpers.RespondJSON(w, models.Health{Health: "ok"})
}
