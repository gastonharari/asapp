package repository

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

const (
	insertNewUserQuery = "INSERT INTO users (user , pwd) VALUES (? , ?);"
	storeTokenQuery    = "UPDATE users SET token=?  WHERE user = ?;"
	getTokenQuery      = "SELECT token FROM users WHERE user = ?;"
	logINUserQuery     = "SELECT id, pwd FROM users WHERE id = ?;"
)

func InitRepo() {
	rows, _ := DB.Query("SELECT id, user ,pwd FROM users")
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
	hashedPWD, _ := bcrypt.GenerateFromPassword([]byte(pwd), 8)

	statement, _ := DB.Prepare(insertNewUserQuery)
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
	r, err := DB.Query(logINUserQuery, id, user)
	if err != nil {
		return 0, err
	}
	r.Scan(&pwd)
	hasedPwd, _ := bcrypt.GenerateFromPassword([]byte(request), 8)
	err = bcrypt.CompareHashAndPassword([]byte(hasedPwd), []byte(request))
	return id, err
}

func StoreToken(user string, token string) error {
	statement, _ := DB.Prepare(storeTokenQuery)
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
	r, err := DB.Query(getTokenQuery, user)
	if err != nil {
		return err
	}
	r.Scan(&token)
	return bcrypt.CompareHashAndPassword([]byte(token), []byte(tokenUser))
}
