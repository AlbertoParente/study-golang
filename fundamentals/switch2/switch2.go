package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good mornig...!")
	case t.Hour() < 18:
		fmt.Println("Good afternoon...!")
	default:
		fmt.Println("Good night...!")
	}
}
