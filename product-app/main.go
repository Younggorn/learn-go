package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

func main() {
	products := []Product{
		{ID: 1, Name: "Keyboard", Price: 1500, Stock: 10},
		{ID: 2, Name: "Mouse", Price: 700, Stock: 20},
		{ID: 3, Name: "Monitor", Price: 5900, Stock: 5},
	}

	jsonData, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}