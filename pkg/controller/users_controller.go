package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/challenge/pkg/helpers"
	"github.com/challenge/pkg/models"
	"github.com/challenge/pkg/usecase"
)

// CreateUser creates a new user
func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var usr models.UserRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&usr)

	if err != nil {
		helpers.RespondBadRequest(w, "bad request, invalid parameters")
		return
	}

	if usr.User == "" {
		fmt.Println(usr.User)
		helpers.RespondBadRequest(w, "the username is required")
		return
	}

	if usr.Password == "" {
		fmt.Println(usr.Password)
		helpers.RespondBadRequest(w, "the password is required")
		return
	}

	userID, err := usecase.CreateNewUser(usr.User, usr.Password)
	if err != nil {
		helpers.RespondError(w, err)
		return
	}
	helpers.RespondJSON(w, models.User{ID: userID})
}
