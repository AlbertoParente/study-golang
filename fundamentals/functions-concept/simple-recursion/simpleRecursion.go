package main

import "fmt"

func factorial(n unit) unit {
	switch {
	case n == 0:
		return 1
	default:
		return m * factorial(n-1)
	}
}

func main() {
	result := factorial(5)
	fmt.Println((result))
}
