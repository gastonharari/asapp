package usecase

import (
	"fmt"

	"github.com/challenge/pkg/repository"
	_ "github.com/mattn/go-sqlite3"
)

func CreateNewUser(user string, password string) (int64, error) {
	fmt.Println(user)
	fmt.Println(password)
	return repository.InserNewUser(user, password)
}
