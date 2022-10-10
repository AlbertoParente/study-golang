package main

import "fmt"

type product struct {
	name     string
	price    float64
	discount float64
}

func (p product) priceWithDiscount() float64 {
	return p.price * (1 - p.discount)
}

func main() {
	var product1 product
	product1 = product{
		name:     "Pen",
		price:    2.59,
		discount: 0.10,
	}
	fmt.Println(product1, product1.priceWithDiscount())

	product2 := product{"Norbook", 4199.00, 0.10}
	fmt.Println(product2.priceWithDiscount())
}
