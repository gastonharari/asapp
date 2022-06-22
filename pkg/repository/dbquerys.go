package repository

import (
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

const (
	insertNewUserQuery = "INSERT INTO users (user , pwd) VALUES (? , ?);"
	storeTokenQuery    = "UPDATE users SET token=?  WHERE user = ?;"
	getTokenQuery      = "SELECT token FROM users WHERE user = ?;"
	logINUserQuery     = "SELECT id, pwd FROM users WHERE user = ?;"
)

var PASSWORD_NOT_CORRECT_ERROR = errors.New("invalid password")

func InitRepo() {
	rows, _ := GetDB().Query("SELECT id, user ,pwd FROM users")
	var id int
	var user string
	var pwd string
	for rows.Next() {
		rows.Scan(&id, &user, &pwd)
		fmt.Println(id, user, pwd)
	}
}

func InserNewUser(user string, pwd string) (int64, error) {
	//Create hash to save in db
	hashedPWD, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	statement, _ := GetDB().Prepare(insertNewUserQuery)
	r, err := statement.Exec(user, hashedPWD)
	if err != nil {
		return 0, err
	}
	if rowsAffected, _ := r.RowsAffected(); rowsAffected == 0 {
		return 0, fmt.Errorf("no rows affected")
	}

	newID, _ := r.LastInsertId()
	return newID, err
}

func LoginUser(user string, request string) (int64, error) {
	var pwd string
	var id int64
	r, err := GetDB().Query(logINUserQuery, user)
	if err != nil {
		return 0, err
	}
	for r.Next() {
		r.Scan(&id, &pwd)
	}
	//hasedPwd, err := bcrypt.GenerateFromPassword([]byte(request), bcrypt.DefaultCost)
	err = bcrypt.CompareHashAndPassword([]byte(pwd), []byte(request))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		err = PASSWORD_NOT_CORRECT_ERROR
	}
	return id, err
}

func StoreToken(user string, token string) error {
	statement, _ := GetDB().Prepare(storeTokenQuery)
	r, err := statement.Exec(token, user)

	if err != nil {
		return err
	}
	if rowsAffected, _ := r.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}
	return nil
}

func ValidateToken(user string, tokenUser string) error {
	var token string
	r, err := GetDB().Query(getTokenQuery, user)
	if err != nil {
		return err
	}
	r.Scan(&token)
	return bcrypt.CompareHashAndPassword([]byte(token), []byte(tokenUser))
}
