package main

import "fmt"

func main() {
	funcSalary := map[string]float64{
		"Alberto Parente":    4000.00,
		"Juliana Cavalcante": 3500.00,
		"Julia":              1200.00,
	}

	funcSalary["Elenildo"] = 1350.00
	delete(funcSalary, "Absent")

	for name, salary := range funcSalary {
		fmt.Println(name, salary)
	}
}
