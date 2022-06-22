package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	db := GetDB()
	if db == nil {
		panic("could not conect to the db, retry later")
	}
}

//Singleton DB
func GetDB() *sql.DB {
	if DB == nil {
		database, _ := sql.Open("sqlite3", "./dbtest.db")
		DB = database
	}
	return DB
}

func InternalInitDB() {
	statement, _ := GetDB().Prepare("CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY, user TEXT, pwd TEXT, token TEXT, UNIQUE(user));")
	statement.Exec()
	// statement, _ = database.Prepare("INSERT INTO users (user , pwd) VALUES (? , ?);")
	// r, _ := statement.Exec("pepe", "hola")
	// fmt.Println(r.LastInsertId())
	rows, _ := GetDB().Query("SELECT id, user ,pwd, token FROM users")
	var id int
	var user string
	var pwd string
	var token string
	for rows.Next() {
		rows.Scan(&id, &user, &pwd, &token)
		fmt.Println(id, user, pwd, token)
	}
}
