package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Println("%d", i)
	}

	fmt.Println("\nMixing...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Even ")
		} else {
			fmt.Print("Odd ")
		}
	}

	for {
		// infinite loop
		fmt.Println("Forever...")
		fmt.Println(time.Second)
	}
}
