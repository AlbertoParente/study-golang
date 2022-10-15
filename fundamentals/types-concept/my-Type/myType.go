package main

import "fmt"

type note float64

func (n note) between(start, end float64) bool {
	return float64(n) >= start && float64(n) <= end
}

func noteToConcept(n note) string {
	if n.between(9.0, 10.0) {
		return "A"
	} else if n.between(7.0, 8.99) {
		return "B"
	} else if n.between(5.0, 7.99) {
		return "C"
	} else if n.between(3.0, 4.99) {
		return "D"
	} else {
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
