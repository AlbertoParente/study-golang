package main

import (
	"fmt"
	"os"
)

func main() {

	showIntroduction()
	showMenu()
	command := readCommand()

	/*
		if command == 1 {
			fmt.Println("Monitoring...")
		} else if command == 2 {
			fmt.Println("Viewing logs...")
		} else if command == 0 {
			fmt.Println("Leaving the program...")
		} else {
			fmt.Println("Invalid command...!")
		}
	*/

	switch command {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Viewing logs...")
	case 0:
		fmt.Println("Leaving the program...")
		os.Exit(0)
	default:
		fmt.Println("Invalid command...!")
		os.Exit(-1)
	}
}

func showIntroduction() {
	name := "Alberto Parente"
	version := 1.1
	fmt.Println("Hello, sr.", name)
	fmt.Println("This program is in version", version)
}

func showMenu() {
	fmt.Println("1- Starting monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit the program")
}

func readCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)
	fmt.Println("The chosen command was", commandRead)

	return commandRead
}
