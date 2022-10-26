package main

import (
	"fmt"
	"time"
)

func speak(person, text string, quantity int) {
	for i := 0; i < quantity; i++ {
		time.Sleep(time.Second)
		fmt.Println("%s: %s (interaction %d)\n", person, text, i+1)
	}
}

func main() {
	// go speak("Juliana", "Hello...", 500)
	// go speak("Alberto", "Hello...", 500)

	// fmt.Println("End!")

	go speak("Juliana", "Hello...", 10)
	speak("ALberto", "Hello...", 5)
}
