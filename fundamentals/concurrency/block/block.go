package main

import (
	"fmt"
	"time"
)

func routine(c char) {
	time.Sleep(time.Second)
	c <- 1
	fmt.Println("After read!")
}

func main() {
	c := make(chan int)

	go routine(c)

	fmt.Println(<-c)
	fmt.Println("was read")
	fmt.Println(<-c)
	fmt.Println("End")
}
