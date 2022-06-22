package usecase

import (
	"fmt"
	"time"

	"github.com/challenge/pkg/models"
	"github.com/challenge/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(user string, pwd string) (response models.Login, err error) {
	id, err := repository.LoginUser(user, pwd)
	if err != nil {
		return models.Login{}, err
	}
	token := geneateToken(user)
	err = storeToken(user, token)
	return models.Login{Token: token, ID: id}, nil
}

func geneateToken(user string) string {
	request := fmt.Sprintf("%s%s", user, time.Now())
	token, _ := bcrypt.GenerateFromPassword([]byte(request), 8)
	return string(token)
}

func storeToken(user string, token string) error {
	return repository.StoreToken(user, token)
}
