package main

import (
	"ecommerce-api/database"
	"ecommerce-api/handler"
	"ecommerce-api/models"

	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")

func main() {

	// connect to DB
	initDb()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route
	e.POST("/login", handler.Loginhandler)

	// Register route
	e.POST("/register", handler.Loginhandler)

	// Product API
	e.GET("/api/v1/product", handler.ProductHandler)
	// r.GET("/product/", handler.ProductHandler)

	// Restricted group
	r := e.Group("/api/v1/user")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &models.MyClaims{},
		SigningKey: JWT_SIGNATURE_KEY,
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", handler.ErrorHandler)

	// User API
	r.GET("/profile", handler.UserHandler)
	r.GET("/orders", handler.OrderHandler)

	// // Cart API
	// r.GET("/cart", handler.CartHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func initDb() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "",
			DB:         "ecommerce",
		}

	connectionString := database.GetConnectionString(config)
	dbErr := database.Connect(connectionString)
	if dbErr != nil {
		panic(dbErr.Error())
	}
}
