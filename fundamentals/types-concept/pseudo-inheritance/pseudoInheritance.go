package main

import "fmt"

type car struct {
	name         string
	currentSpeed int
}

type ferrari struct {
	car
	activeTurbo bool
}

func main() {
	f := ferrari{}
	f.name = "F40"
	f.currentSpeed = 0
	f.activeTurbo = true

	fmt.Println("Does the Ferrari %s have the turbo on? %w\n", f.name, f.activeTurbo)
	fmt.Println(f)
}
