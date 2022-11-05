package main

import (
	"fmt"
	"time"
)

func speak(person string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s speak: %d", person, i)
		}
	}()
	return c
}

func beatween(input1, input2, input3, input4 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			case s := <-input3:
				c <- s
			case s := <-input4:
				c <- s
			}
		}
	}()
	return c
}

func main() {
	c := beatween(speak("Alberto Parente"), speak("Juliana Cvalcante"),
		speak("Julia Parente"), speak("Alberto"))
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
}
