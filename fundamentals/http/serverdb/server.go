package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users/", UserHandler)
	log.Println("Starting...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
