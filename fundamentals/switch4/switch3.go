package main

import "fmt"

func typee(i interface{}) string {
	switch i.(type) {
	case int:
		return "Integer"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case func():
		return "Function"
	default:
		return "Not exists"
	}
}

func main() {
	fmt.Println(type(2, 3))
    fmt.Println(type(1))
    fmt.Println(type("Hello"))
    fmt.Println(type(func() {}))
    fmt.Println(type(time.Now()))
}
