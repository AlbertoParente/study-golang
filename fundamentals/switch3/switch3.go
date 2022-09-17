package main

import "fmt"

func noteToConcept(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(noteToConcept(10))
	fmt.Println(noteToConcept(9.8))
	fmt.Println(noteToConcept(6.9))
	fmt.Println(noteToConcept(2.1))
	fmt.Println(noteToConcept(1))
}
