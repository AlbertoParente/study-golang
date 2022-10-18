package main

import "fmt"

type sports interface {
	turnOnTurbo()
}

type luxurious interface {
	makeGoal()
}

type sportsLuxurious interface {
	sports
	luxurious
}

type bmw7 struct{}

func (b bmw7) turnOnTurbo() {
	fmt.Println("Turbo...")
}

func (b bmw7) makeGoal() {
	fmt.Println("Hey....")
}

func main() {
	var b sportsLuxurious = bmw7{}
	b.turnOnTurbo()
	b.makeGoal()
}
