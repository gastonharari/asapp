package usecase

import (
	"fmt"

	"github.com/challenge/pkg/repository"
	_ "github.com/mattn/go-sqlite3"
)

func CreateNewUser(user string, password string) int64 {
	fmt.Println(user)
	fmt.Println(password)
	db := repository.InitDb()
	index := repository.InserNewUser(db, user, password)
	return index
}
