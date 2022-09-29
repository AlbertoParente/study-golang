package main

import "fmt"

func mult(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func exec(function func(int, int) int, p1, p2 int) int {
	return function(p1, p2)
}

func exec1(function func(int, int, int) int, p1, p2, p3 int) int {
	return function(p1, p2, p3)
}

func main() {
	result := exec(mult, 3, 4)
	result1 := exec1(div, 3, 4, 5)
	fmt.Println(result)
	fmt.Println(result1)
}
