package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-drive/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO USERS(name) VALUES(?)")
	stmt.Exec("Alberto")
	stmt.Exec("Juliana")

	res, _ := stmt.Exec("Julia")

	id, _ := res.LastInsertId()
	fmt.Println(id)

	lines, _ := res.RowsAffected()
	fmt.Println(lines)

}
