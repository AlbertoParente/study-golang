package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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
			printLog()
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

	if resp.StatusCode == 200 {
		fmt.Println("The Site:", site, "was successfully loaded!")
	} else {
		fmt.Println("The Site:", site, "did not load successfully! Status code:", resp.StatusCode)
	}
}

func getSitesFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("An error has occurred:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func registerLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("An error has occurred:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site +
		" - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLog() {

	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("An error has occurred:", err)
	}

	fmt.Println(string(file))
}
