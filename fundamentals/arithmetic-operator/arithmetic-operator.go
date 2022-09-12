package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Sum =>", a+b)
	fmt.Println("Subtraction =>", a-b)
	fmt.Println("Multiplication =>", a*b)
	fmt.Println("Division =>", a/b)
	fmt.Println("Module =>", a%b)

	// bitwise
	fmt.Println("AND =>", a&b)
	fmt.Println("OR =>", a|b)
	fmt.Println("XOR =>", a^b)

	c := 3.0
	d := 2.0

	fmt.Println("Is bigger =>", math.Max(float64(a), float64(b)))
	fmt.Println("Is smaller =>", math.Min(c, d))
	fmt.Println("Exponential =>", math.Pow(c, d))
}
