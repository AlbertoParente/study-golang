package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi((sid))

	switch {
	case r.Method == "GET" && id > 0:
		userPorID(w, r, id)
	case r.Method == "GET":
		allUser(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Sorry... :(")
	}
}

func userPorID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:123456@/studygo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var u User
	db.QueryRow("SELECT id, name, surname, email FROM USERS WHERE id = ?", id).
		Scan(&u.ID, &u.Name, &u.Surname, &u.Email)

	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func allUser(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:123456@/studygo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("SELECT id, name FROM USERS")
	defer db.Close()

	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Email)
		users = append(users)
	}

	json, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
