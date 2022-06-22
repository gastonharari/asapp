package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

const (
	insertNewUserQuery = "INSERT INTO users (user , pwd) VALUES (? , ?);"
)

func InitRepo(db *sql.DB) {
	rows, _ := db.Query("SELECT id, user ,pwd FROM users")
	var id int
	var user string
	var pwd string
	for rows.Next() {
		rows.Scan(&id, &user, &pwd)
		fmt.Println(id, user, pwd)
	}
}

func InserNewUser(db *sql.DB, user string, pwd string) int64 {
	//Create hash to save in db
	hashedPWD, _ := bcrypt.GenerateFromPassword([]byte(pwd), 8)

	statement, _ := db.Prepare(insertNewUserQuery)
	r, _ := statement.Exec(user, hashedPWD)
	newID, _ := r.LastInsertId()
	return newID
}
