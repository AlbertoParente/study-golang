package main

import "fmt"

func purchase(t1, t2 bool) (bool, bool, bool) {
	Tv50 := t1 && t2
	Tv30 := t1 != t2
	IceCream := t1 || t2

	return Tv30, Tv50, IceCream
}

func main() {
	Tv50, Tv32, IceCream := purchase(true, true)
	fmt.Println("Tv50: %t, Tv32: %t, IceCream: %t, Healthy: %t",
		Tv50, Tv32, IceCream, !IceCream)
}
