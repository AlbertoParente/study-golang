package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomNumber() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := RandomNumber(); i > 5 {
		fmt.Println("Winner!!!")
		fmt.Println(i)
	} else {
		fmt.Println("Loser")
	}
}
