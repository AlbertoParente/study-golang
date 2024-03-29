package main

import "fmt"

type printable interface {
	toSring() string
}

type person struct {
	name    string
	surname string
}

type product struct {
	name  string
	price float64
}

func (p person) toString() string {
	return p.name + " " + p.surname
}

func (p product) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.name, p.price)
}

func printt(x printable) {
	fmt.Println(x.toSring())
}

func main() {
	var something printable = person{"Alberto", "Parente"}
	fmt.Println(something.toSring())
	printt(something)

	something = product{"Notbook", 1499.99}
	fmt.Println(something.toSring())
	printt(something)

	p2 := product{"Notbook", 2000.00}
	printt(p2)
}
