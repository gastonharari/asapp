package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InitDb() *sql.DB {
	database, _ := sql.Open("sqlite3", "./dbtest.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY, user TEXT, pwd TEXT, token TEXT, UNIQUE(user));")
	statement.Exec()
	// statement, _ = database.Prepare("INSERT INTO users (user , pwd) VALUES (? , ?);")
	// r, _ := statement.Exec("pepe", "hola")
	// fmt.Println(r.LastInsertId())
	rows, _ := database.Query("SELECT id, user ,pwd, token FROM users")
	var id int
	var user string
	var pwd string
	var token string
	for rows.Next() {
		rows.Scan(&id, &user, &pwd, &token)
		fmt.Println(id, user, pwd, token)
	}
	return database
}
