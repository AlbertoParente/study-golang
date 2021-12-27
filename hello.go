package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	showIntroduction()
	showMenu()

	name, age := getDados()
	fmt.Println(name, age)

	command := readCommand()

	switch command {
	case 1:
		startingMonitoring()
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

func getDados() (string, int) {
	name := "Alberto Parente"
	age := 24
	return name, age
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

func startingMonitoring() {
	fmt.Println("Monitoring...")
	site := "https://www.innovaro.com.br"
	resp, _ := http.Get(site)
	fmt.Println(resp)
}
