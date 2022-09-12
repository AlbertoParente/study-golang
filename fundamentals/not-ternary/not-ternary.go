package main

import "fmt"

// Not exists ternary operator in golang

func getResult(note float64) string {
	if note >= 6 {
		return "Approved"
	}
	return "Reproved"

	// Not possible this
	// return note >= 6 ? "Approved" : "Reproved"
}

func main() {
	fmt.Println(getResult(6.2))
}
