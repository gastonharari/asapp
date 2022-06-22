package usecase

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/challenge/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(user string, pwd string) (token string, err error) {
	db := repository.InitDb()
	validLogin, err := repository.LoginUser(db, user, pwd)
	if err != nil {
		return "", nil
	}
	fmt.Println(validLogin)
	token = geneateToken(user)
	err = storeToken(db, user, token)
	return token, err
}

func geneateToken(user string) string {
	request := fmt.Sprintf("%s%s", user, time.Now())
	token, _ := bcrypt.GenerateFromPassword([]byte(request), 8)
	return string(token)
}

func storeToken(db *sql.DB, user string, token string) error {
	return repository.StoreToken(db, user, token)
}
