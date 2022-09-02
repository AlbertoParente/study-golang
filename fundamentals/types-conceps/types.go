package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal integer is:", reflect.TypeOf(32000))

	var b byte = 255
	fmt.Println("The byte is:", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("The value max of int is:", i1)
	fmt.Println("The type of i1 is:", reflect.TypeOf(i1))

	var i2 rune = 'a'
	fmt.Println("The rune is ", reflect.TypeOf(i2))
	fmt.Println(i2)

	var x float32 = 49.99
	fmt.Println("O type of x is:", reflect.TypeOf(x))
	fmt.Println("The type of literal is:", reflect.TypeOf(49.99))

	bo := true
	fmt.Println("O type of bo is:", reflect.TypeOf(bo))
	fmt.Println(!bo)

	s1 := "Hello"
	fmt.Println(s1 + "!")
	fmt.Println("The size of string is:", len(s1))

	s2 := `Hello
    my
    name
    is
    Alberto Parente`
	fmt.Println(s2 + "!")
	fmt.Println("The size of string is:", len(s1))

	// Char not exists in linguage Go
	char := 'a'
	fmt.Println("The type of char is:", reflect.TypeOf(char))
	fmt.Println(char) // return value is unicode

}
