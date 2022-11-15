package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-drive/mysql"
)

type car struct {
	id    int
	model string
	mark  string
}

type motorcycle struct {
	id    int
	model string
	mark  string
}

type user struct {
	id      int
	name    string
	surname string
	address string
	email   string
}

func main() {
	db, err := sql.Open("mysql", "root:123456@/TestGo")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _ := db.Query("SELECT id, model, mark FROM CAR WHERE ID > ?", 3)
	rows1, _ := db.Query("SELECT id, model, mark FROM MOTORCYCLE WHERE ID > ?", 3)
	rows2, _ := db.Query("SELECT id, name, surname, address, email FROM USERS WHERE ID > ?", 3)
	defer rows.Close()
	defer rows1.Close()
	defer rows2.Close()

	for rows.Next() {
		var i car
		rows.Scan(&i.id, &i.model)
		fmt.Println(i)
	}

	for rows1.Next() {
		var j motorcycle
		rows.Scan(&j.id, &j.model)
		fmt.Println(j)
	}

	for rows2.Next() {
		var l user
		rows.Scan(&l.id, &l.name)
		fmt.Println(l)
	}
}
