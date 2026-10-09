package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

func validateProduct(product Product) error {
	if product.Name == "" {
		return fmt.Errorf("product name is required")
	}
	if product.Price <= 0 {
		return fmt.Errorf("product price must be greater than zero")
	}
	if product.Stock < 0 {
		return fmt.Errorf("product stock cannot be negative")
	}
	return nil
}

func findProductByID(id int, products []Product) (*Product, error) {
	for i := range products {
		if products[i].ID == id {
			return &products[i], nil
		}
	}

	return nil, fmt.Errorf("product with ID %d not found", id)
}

func main() {

	products := []Product{
		{ID: 1, Name: "Keyboard", Price: 1500, Stock: 10},
		{ID: 2, Name: "Mouse", Price: 700, Stock: 20},
		{ID: 3, Name: "Monitor", Price: 5900, Stock: 5},
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go API")
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:

			w.Header().Set("Content-Type", "application/json")

			err := json.NewEncoder(w).Encode(products)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		case http.MethodPost:
			var product Product

			err := json.NewDecoder(r.Body).Decode(&product)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if err := validateProduct(product); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			product.ID = len(products) + 1
			products = append(products, product)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			err = json.NewEncoder(w).Encode(product)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}

	})

	http.HandleFunc("GET /products/{id}", func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")

		productID, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		product, err := findProductByID(productID, products)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("PUT /products/{id}", func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")

		productID, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		product, err := findProductByID(productID, products)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		var updatedProduct Product

		err = json.NewDecoder(r.Body).Decode(&updatedProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := validateProduct(updatedProduct); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		product.Name = updatedProduct.Name
		product.Price = updatedProduct.Price
		product.Stock = updatedProduct.Stock

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(product); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Server running on port 3000")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
