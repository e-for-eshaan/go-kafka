package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"web-server/database"
	"web-server/kafkaqueue"

	"github.com/lib/pq"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	rows, err := database.DB.Query("SELECT id, name, mobile, latitude, longitude, created_at, updated_at FROM users")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error fetching users")
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Mobile, &user.Latitude, &user.Longitude, &user.CreatedAt, &user.UpdatedAt); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Error scanning users")
			return
		}
		users = append(users, user)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}


func GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Fetch products from the database
	rows, err := database.DB.Query("SELECT product_id, product_name, product_description, product_images, product_price, compressed_product_images, created_at, updated_at FROM products")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error fetching products from the database")
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ProductID, &product.ProductName, &product.ProductDescription,
			pq.Array(&product.ProductImages), &product.ProductPrice, pq.Array(&product.CompressedProductImgs),
			&product.CreatedAt, &product.UpdatedAt); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Error scanning products")
			return
		}
		products = append(products, product)
	}

	// Set the response Content-Type to "application/json" and send the products as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}


func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body into a Product struct
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid request payload")
		return
	}

	query := `INSERT INTO products (product_name, product_description, product_images, product_price, compressed_product_images, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING product_id`
	err = database.DB.QueryRow(query, product.ProductName, product.ProductDescription, pq.Array(product.ProductImages),
		product.ProductPrice, pq.Array(product.CompressedProductImgs), time.Now(), time.Now()).Scan(&product.ProductID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error inserting product into the database")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
    kafkaqueue.PublishToKafka(product.ProductID)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid request payload")
		return
	}

	query := `INSERT INTO users (name, mobile, latitude, longitude, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err = database.DB.QueryRow(query, user.Name, user.Mobile, user.Latitude, user.Longitude, time.Now(), time.Now()).Scan(&user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error inserting user into the database")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}


func GetImageFromQueue(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	productID := kafkaqueue.ConsumeFromKafka()

	// Fetch the product from the database based on the given product ID
	row := database.DB.QueryRow("SELECT product_id, product_name, product_description, product_images, product_price, compressed_product_images, created_at, updated_at FROM products WHERE product_id = $1", productID)

	var product Product
	if err := row.Scan(&product.ProductID, &product.ProductName, &product.ProductDescription,
		pq.Array(&product.ProductImages), &product.ProductPrice, pq.Array(&product.CompressedProductImgs),
		&product.CreatedAt, &product.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "Product not found")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Error fetching product from the database")
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
