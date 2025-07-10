package main

import (
	"log"
	"os"

	"bank-it/db"
	"bank-it/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	//Establish DB connection and create tables if not exists
	database := db.Connect()
	db.RunMigrations(database)

	//Define API routers
	router := gin.Default()

	//Register the handlers
	handlers.RegisterAccountRoutes(router, database)
	handlers.RegisterTransactionRoutes(router, database)

	//Set PORT from .env or defaults to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//Start the server
	router.Run(":" + port)
	log.Println("Server running on :" + port)
}
