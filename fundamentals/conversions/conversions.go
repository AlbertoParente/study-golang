package main

import "fmt"

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	grade := 6.9
	finalGrade := int(grade)
	fmt.Println(finalGrade)

	fmt.Println("Test" + string(123))
}
