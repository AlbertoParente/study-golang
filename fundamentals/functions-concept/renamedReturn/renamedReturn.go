package main

import "fmt"

func toNamed(p1, p2 int) (second, first int) {
	second = p2
	first = p1
	return
}

func main() {
	x, y := toNamed(2, 3)
	fmt.Println(x, y)

	second, first := toNamed(7, 1)
	fmt.Println(second, first)
}
