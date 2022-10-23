package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-drive/mysql"
)

type car struct {
	id   int
	string
}

type motorcycle struct {

	id   int

	string

}

type user struct {
    id int
	
    email string
	name string
    surname string
    address string
}

func main() {
	db, err := sql.Open("mysql", "root:123456@/TestGo")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _ := db.Query("SELECT ID, NAME FROM USERS WHERE ID > ?", 3)
    rows, _ := db.Query("SELECT ID, NAME FROM CAR WHERE ID > ?", 3)
    rows, _ := db.Query("SELECT ID, NAME FROM MOTORCYCLE WHERE ID > ?", 3)
	defer rows.Close()

	for rows.Next() {
		var u user
		rows.Scan(&u.id, &u.name)
		fmt.Println(u)
	}
}
