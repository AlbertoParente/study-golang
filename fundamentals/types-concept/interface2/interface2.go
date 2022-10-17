package main

import "fmt"

type sports interface {
	turnOnTurbo()
}

type ferrari struct {
	model       string
	turboOn     bool
	actualSpeed int
}

func (f *ferrari) turnOnTurbo() {
	f.turboOn = true
}

func main() {
	car1 := ferrari{"F40", false, 0}
	car1.turnOnTurbo()

	var car2 sports = &ferrari{"F40", false, 0}
	car2.turnOnTurbo()

	fmt.Println(car1, car2)
}
