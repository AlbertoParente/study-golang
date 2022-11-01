package main

import (
	"fmt"
	"html"
)

func toForward(origin <-chan string, destin chan string) {
	for {
		destin <- <-origin
	}
}

func join(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go toForward(input1, c)
	go toForward(input2, c)
	return c
}

func main() {
	c := join(
		html.Title("https://www.google.com.br", "https://www;innovaro.com.br"),
		html.Title("www.amazon.com", "https://www.youtube.com.br"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
