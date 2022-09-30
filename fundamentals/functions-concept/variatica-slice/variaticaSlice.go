package main

import "fmt"

func showApproved(approved ...string) {
	fmt.Println("List of approved")
	for i, approved := range approved {
		fmt.Println("%d) %s\n", i+1, approved)
	}
}

func main() {
	approved := []string{"Alberto", "Juliana", "Julia", "Elenildo"}
	showApproved(approved...)
}
