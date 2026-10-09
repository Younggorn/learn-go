package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"context"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type UpdateProductRequest struct {
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

func validateUpdateProduct(product UpdateProductRequest) error {
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

	config, err := loadConfig()
	if err != nil {
		fmt.Println("Configuration error:", err)
		return
	}

	db, err := connectDB(config)
	if err != nil {
		fmt.Println("Database error:", err)
		return
	}

	defer db.Close()

	fmt.Println("PostgreSQL connected successfully")

	rows, err := db.Query(
		context.Background(),
		"SELECT id, name, price, stock FROM products ORDER BY id",
	)

	if err != nil {
		fmt.Println("Query error:", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var product Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Stock,
		)

		if err != nil {
			fmt.Println("Scan error:", err)
			return
		}

		fmt.Println(product)
	}

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

		var updatedProduct UpdateProductRequest

		err = json.NewDecoder(r.Body).Decode(&updatedProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := validateUpdateProduct(updatedProduct); err != nil {
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

	http.HandleFunc("DELETE /products/{id}", func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")

		productID, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		deleted := false

		// Remove the product from the slice
		for i, p := range products {
			if p.ID == productID {
				products = append(products[:i], products[i+1:]...)
				break
			}
		}

		if !deleted {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})

	fmt.Println("Server running on port 3000")

	if err := http.ListenAndServe("127.0.0.1:3000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
