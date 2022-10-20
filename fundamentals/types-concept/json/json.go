package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []dtring `json:"Tags"`
}

func main() {
	p1 := product{1, "Celular", 1999.99, []string{"Promotion", "Electronic"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	var p2 product
	jsonString := `{"id":2,"name":"Pencil","price":8.90,"tags":["stationners", "Importing"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
}
