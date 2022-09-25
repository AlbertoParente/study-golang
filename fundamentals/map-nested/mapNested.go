package main

import "fmt"

func main() {
	funcByLetters := map[string]map[string]float64{
		"J": {
			"João Alberto":       78445.78,
			"Juliana Cavalcante": 7894.78,
		},
		"G": {
			"Gabriella": 1234.58,
		},
		"E": {
			"Elenildo": 1634.56,
		},
	}

	delete(funcByLetters, "E")

	for letter, funcs := range funcByLetters {
		fmt.Println(letter, funcs)
	}
}
