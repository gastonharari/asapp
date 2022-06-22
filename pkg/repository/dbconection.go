package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InitDb() *sql.DB {
	database, _ := sql.Open("sqlite3", "./dbtest.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY, user TEXT, pwd TEXT)")
	statement.Exec()
	/*statement, _ = database.Prepare("INSERT INTO users (user , pwd) VALUES (? , ?);")
	r, _ := statement.Exec("pepe", "hola")
	fmt.Println(r.LastInsertId())*/
	rows, _ := database.Query("SELECT id, user ,pwd FROM users")
	var id int
	var user string
	var pwd string
	for rows.Next() {
		rows.Scan(&id, &user, &pwd)
		fmt.Println(id, user, pwd)
	}
	return database
}
