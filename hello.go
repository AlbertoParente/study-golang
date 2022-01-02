package main

import (
	"fmt"
	"net/http"
	"os"
)

const numberOfMonitoring = 3
const delay = 5

func main() {

	showIntroduction()

	for {
		showMenu()

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
	fmt.Println("")

	return commandRead
}

func startingMonitoring() {
	fmt.Println("Monitoring...")

	sites := getSitesFile()

	for i := 0; i < numberOfMonitoring; i++ {
		for i, site := range sites {
			fmt.Println(i+1, site)
			siteTests(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func siteTests(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("An error has occurred:", err)
	}
