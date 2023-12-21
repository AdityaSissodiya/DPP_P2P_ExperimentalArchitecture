package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var dbProducts *sql.DB

func init() {
	var err error
	dbProducts, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	_, err = dbProducts.Exec(`CREATE TABLE IF NOT EXISTS Products (
		ProductID INTEGER PRIMARY KEY,
		Name TEXT NOT NULL,
		Manufacturer TEXT,
		ReleaseDate DATE,
		Description TEXT
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// Insert sample data
	_, err = dbProducts.Exec(`INSERT INTO Products (Name, Manufacturer, ReleaseDate, Description) VALUES
		('Laptop Battery', 'ABC Electronics', '2023-01-01', 'High-performance laptop battery'),
		('Smartphone Battery', 'XYZ Tech', '2023-02-15', 'Long-lasting battery for smartphones')`)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/products", getProducts).Methods("GET")

	fmt.Println("Server A (Products) is running on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := dbProducts.Query("SELECT ProductID, Name, Manufacturer, ReleaseDate, Description FROM Products")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ProductID, &product.Name, &product.Manufacturer, &product.ReleaseDate, &product.Description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// Product represents the structure of the Product table
type Product struct {
	ProductID    int    `json:"productID"`
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	ReleaseDate  string `json:"releaseDate"`
	Description  string `json:"description"`
}
