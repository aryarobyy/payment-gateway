package main

import (
	"fmt"
	"log"
	"os"

	"payment-gateway/app"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: .env file not found or could not be loaded")
	}

	application := app.NewApp()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s with CORS enabled", port)

	if err := application.Run("0.0.0.0:" + port); err != nil {
		fmt.Println("FATAL ERROR: Server failed to run:", err)
		os.Exit(1)
	}
}
