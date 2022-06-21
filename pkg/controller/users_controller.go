package controller

import (
	"fmt"
	"net/http"

	"github.com/challenge/pkg/helpers"
	"github.com/challenge/pkg/models"
	"github.com/challenge/pkg/usecase"
)

// CreateUser creates a new user
func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		//TODO: Handle error
		return
	}

	usr := r.PostForm.Get("username")
	pwd := r.PostForm.Get("password")

	fmt.Println(usr)

	if usr == "" {
		fmt.Println(usr)
		helpers.RespondBadRequest(w, "the username is required")
		return
	}

	if pwd == "" {
		fmt.Println(pwd)
		helpers.RespondBadRequest(w, "the password is required")
		return
	}

	userID := usecase.CreateNewUser(usr, pwd)
	helpers.RespondJSON(w, models.User{ID: userID})
}
