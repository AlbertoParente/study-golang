package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-drive/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/TestGo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO USERS(Id, Name) VALUES(?, ?)")

	stmt.Exec(2000, "Alberto")
	stmt.Exec(2001, "Juliana")
	_, err = stmt.Exec(1, "Julia")

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
