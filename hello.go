package main

import (
	"fmt"
	"net/http"
	"os"
)

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

	return commandRead
}

func startingMonitoring() {
	fmt.Println("Monitoring...")
	sites := []string{
		"https://www.innovaro.com.br",
		"http://sync.innovaro.com.br/sync/",
		"http://licenca.innovaro.com.br/allylicenca/",
		"http://app.allysistemas.com.br/sync/",
	}

	for i, site := range sites {
		fmt.Println(i, site)
	}
}

func siteTests(site string) {

    resp, _ := http.Get(site)

    if resp.StatusCode == 200 {
        fmt.Println("The Site:", site, "was successfully loaded!")
    } else {
        fmt.Println("The Site:", site, "did not load successfully! Status code:", resp.StatusCode)
    }
}
