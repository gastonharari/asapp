package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	_, err := GetDB()
	if err != nil {
		panic("could not conect to the db, retry later")
	}
}

//Singleton DB
func GetDB() (*sql.DB, error) {
	if DB == nil {
		database, err := sql.Open("sqlite3", "./dbtest.db")
		if err != nil {
			return nil, err
		}
		DB = database
	}
	return DB, nil
}

func InternalInitDB(db *sql.DB) {
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY, user TEXT, pwd TEXT, token TEXT, UNIQUE(user));")
	statement.Exec()
	// statement, _ = database.Prepare("INSERT INTO users (user , pwd) VALUES (? , ?);")
	// r, _ := statement.Exec("pepe", "hola")
	// fmt.Println(r.LastInsertId())
	rows, _ := db.Query("SELECT id, user ,pwd, token FROM users")
	var id int
	var user string
	var pwd string
	var token string
	for rows.Next() {
		rows.Scan(&id, &user, &pwd, &token)
		fmt.Println(id, user, pwd, token)
	}
}
