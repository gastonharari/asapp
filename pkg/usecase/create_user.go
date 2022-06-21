package usecase

import "fmt"

func CreateNewUser(user string, password string) int64 {
	fmt.Println(user)
	fmt.Println(password)
	return 0
}
