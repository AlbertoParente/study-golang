package main

import "fmt"

func main() {
	numbers := [...]int{1, 2, 3, 4, 5}

	for i, number := range numbers {
		fmt.Println("%d) %d\n", i+1, number)
	}

	for _, num := range numbers {
		fmt.Println(num)
	}

}
