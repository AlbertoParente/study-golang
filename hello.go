package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Alberto Parente"
	age := 29
	version := 1.0
	fmt.Println("Hello world, my name is", name)
	fmt.Println("My age is,", age)
	fmt.Println("This version is,", version)

	name2 := "Alberto Parente"
	age2 := 29
	version2 := 1.0
	fmt.Println("Hello world, my name is", name2, "my age is", age2)
	fmt.Println("This version is,", version2)
	fmt.Println("The type of variable name is:", reflect.TypeOf((name2)))

	fmt.Println("1- Starting monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit the program")

	var command int
	fmt.Scan(&command)
	fmt.Println("The memory address is:", &command)
	fmt.Println("The command chosen was:", command)
	
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
	}
}

func showIntroduction() {
    name := "Alberto Parente"
    version := 1.1
    fmt.Println("Hello, sr.", name)
    fmt.Println("This program is in version", version)
}

func readCommand() int {
    var commandRead int
    fmt.Scan(&commandRead)
    fmt.Println("The chosen command was", commandRead)

    return commandRead
}
