package main

import (
	"fmt"
	"time"
)

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
	fmt.Println(typee(2.3))
	fmt.Println(typee(1))
	fmt.Println(typee("Hello"))
	fmt.Println(typee(func() {}))
	fmt.Println(typee(time.Now()))
}
