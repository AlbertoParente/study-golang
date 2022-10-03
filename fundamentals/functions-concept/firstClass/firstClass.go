package main

import "fmt"

var summ = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(summ(2, 3))
	sub := func(a, b int) int {
		return a - b
	}

	div := func(a, b int) int {
		return a / b
	}

	mult := func(a, b int) int {
		return a * b
	}

	fmt.Println(sub(2, 3))
	fmt.Println(div(10, 2))
	fmt.Println(mult(5, 3))
}
