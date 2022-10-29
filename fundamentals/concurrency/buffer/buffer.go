package main

import (
	"fmt"
	"time"
)

func routine(ch chan int) {
	fmt.Println("Performed!")
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	ch <- 6
}

func main() {
	ch := make(chan int, 3)
	go routine(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
