package main

import "fmt"

func printResult(note float64) {
	if note >= 7 {
		fmt.Println("Approved:", note)
	} else {
		fmt.Println("Reproved:", note)
	}
}

func main() {
	printResult(7.3)
	printResult(5.1)
	printResult(7)
}
