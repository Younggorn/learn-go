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
	jsonData := `{
	"id": 10,
	"name": "Gaming Mouse",
	"price": 1290,
	"stock": 15
}`

	var product Product
	err := json.Unmarshal([]byte(jsonData), &product)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("ID: %d\n", product.ID)
	fmt.Printf("Name: %s\n", product.Name)
	fmt.Printf("Price: %.2f\n", product.Price)
	fmt.Printf("Stock: %d\n", product.Stock)


	
}