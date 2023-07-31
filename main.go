package main

import (
	"log"
	"net/http"
	"web-server/api"
	"web-server/database"

	"github.com/rs/cors"
)

func main() {
	corsHandler := cors.Default()

	if err := database.InitDB(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	defer database.DB.Close()

	http.Handle("/users", corsHandler.Handler(http.HandlerFunc(api.GetUsers)))
	http.Handle("/products", corsHandler.Handler(http.HandlerFunc(api.GetProducts)))
	http.Handle("/products-post", corsHandler.Handler(http.HandlerFunc(api.CreateProduct)))
	http.Handle("/users-post", corsHandler.Handler(http.HandlerFunc(api.CreateUser)))
	http.Handle("/get-from-kafa", corsHandler.Handler(http.HandlerFunc(api.GetImageFromQueue)))

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
