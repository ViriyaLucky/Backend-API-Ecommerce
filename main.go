package main

import (
	"ecommerce-api/database"
	"ecommerce-api/handler"
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {

	// connect to DB
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "root",
			DB:         "learning",
		}

	connectionString := database.GetConnectionString(config)
	dbErr := database.Connect(connectionString)
	if dbErr != nil {
		panic(dbErr.Error())
	}

	mux := http.NewServeMux()

	// User API
	mux.HandleFunc("/api/user/login", handler.Loginhandler)
	mux.HandleFunc("/api/user/register", handler.RegisterHandler)
	mux.HandleFunc("/api/user", handler.UserHandler)

	// Product API
	mux.HandleFunc("/api/product", handler.ProductHandler)

	// Cart API
	mux.HandleFunc("/api/cart", handler.CartHandler)

	// Other route
	mux.HandleFunc("/", handler.ErrorHandler)
	log.Println("Starting on port 8080")

	// Error log
	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)
}
