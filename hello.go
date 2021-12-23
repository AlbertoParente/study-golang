package main

import "fmt"
import "reflect"

func main() {
	var name string = "Alberto Parente"
	var age int = 29
	var version float32 = 1.0
	fmt.Println("Hello world, my name is", name)
	fmt.Println("My age is,", age)
	fmt.Println("This version is,", version)

	var name2 = "Alberto Parente"
	var age2 = 29
	var version2 float32 = 1.0
	fmt.Println("Hello world, my name is", name2, "my age is", age2)
	fmt.Println("This version is,", version2)

	fmt.Println("The type of variable name is:", reflect.TypeOf((name2)))
}
