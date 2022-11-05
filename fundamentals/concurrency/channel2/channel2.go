package main

import (
	"fmt"
	"time"
)

func twoThreeFourTime(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base

	time.Sleep(100 / time.Second)
	c <- 5 * base
}

func main() {
	c := make(chan int)
	go twoThreeFourTime(2, make(chan int))
	fmt.Println("1")

	a, b := <-c, <-c
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<-c)
}
