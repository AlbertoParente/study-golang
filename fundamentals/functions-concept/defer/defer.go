package main

import "fmt"

func getApprovedValue(number int) int {
	defer fmt.Println("End!")
	if number >= 5000 {
		fmt.Println("Value is tall...")
		return 5000
	}
	fmt.Println("Value is small...")
	return number
}

func main() {
	fmt.Println(getApprovedValue(6000))
	fmt.Println(getApprovedValue(5000))
	fmt.Println(getApprovedValue(3000))
	fmt.Println(getApprovedValue(1000))
}
