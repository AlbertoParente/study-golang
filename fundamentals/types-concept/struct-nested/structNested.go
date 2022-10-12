package main

import "fmt"

type item struct {
	productID int
	qtd       int
	price     float64
}

type order struct {
	userID int
	itens  []item
}

func (o order) totalValue() float64 {
	total := 0.0
	for _, item := range o.itens {
		total += item.price * float64(item.qtd)
	}
	return total
}

func main() {
	order := order{
		userID: 1,
		itens: []item{
			item{1, 2, 12.10},
			item{2, 1, 33.99},
			item{3, 100, 13.59},
		},
	}

	fmt.Println("The value total is: R$ %.2f", order.totalValue())
}
