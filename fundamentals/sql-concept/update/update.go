package main

import (
	"database/sql"

	_ "github.com/go-sql-drive/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/TestGo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// UPDATE
	stmt, _ := db.Prepare("UPDATE USERS SET NAME = ? WHERE ID = ?")
	stmt.Exec("Uóxiton Istive", 1)
	stmt.Exec("Sheristone Uasleska", 2)

	// DELETE
	stmt2, _ := db.Prepare("UPDATE USERS SET NAME = ? WHERE ID = ?")
	stmt2.Exec(3)
	stmt2.Exec(2)
	stmt2.Exec(4)
}
