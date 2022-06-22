package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/challenge/pkg/helpers"
	"github.com/challenge/pkg/models"
	"github.com/challenge/pkg/usecase"
)

// Login authenticates a user and returns a token
func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	var request models.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		helpers.RespondBadRequest(w, "bad request, invalid parameters")
		return
	}

	if request.User == "" {
		fmt.Println(request.User)
		helpers.RespondBadRequest(w, "the username is required")
		return
	}

	if request.Password == "" {
		fmt.Println(request.Password)
		helpers.RespondBadRequest(w, "the password is required")
		return
	}

	response, err := usecase.LoginUser(request.User, request.Password)

	if err != nil {
		helpers.RespondError(w, err)
		return
	}

	helpers.RespondJSON(w, response)
}
