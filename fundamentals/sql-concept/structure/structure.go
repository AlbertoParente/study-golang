package main

import (
	"database/sql"

	_ "github.com/go-sql-drive/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:123456@/TestGo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	exec(db, "CREATE DATABASE IF NOT EXISTS TESTGO")
	exec(db, "USE TESTGO")
	exec(db, "DROP TABLE IF NOT EXISTS USERS")
	exec(db, `CRETE TABLE USERS (
        id INTEGER AUTO_INCREMENT,
        name VARCGAR(80),
        PRIMARY KEY (id)
    )`)

}
