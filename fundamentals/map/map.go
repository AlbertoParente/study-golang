package main

import "fmt"

func main() {
	// var approveds map[int]string
	approveds := make(map[int]string)
	approveds[215616516516516] = "Alberto Parente"
	approveds[618619613216514] = "Juliana"
	approveds[874981352135468] = "Julia"
	fmt.Println(approveds)

	for cpf, name := range approveds {
		fmt.Println("%s (CPF: %d)\n", name, cpf)
	}

	fmt.Println(approveds[874981352135468])
	delete(approveds, 874981352135468)
	fmt.Println(approveds[874981352135468])
}
