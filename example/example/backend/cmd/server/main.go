package main

import (
	"log"
	"os"

	"backend/config"
	"backend/internal/api"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	if err := config.LoadConfig(); err != nil {
		log.Fatal("Failed to load config:", err)
	}
	cfg := config.GetConfig()

	router := gin.Default()
	router.Use(api.CORSMiddleware())
	api.SetupRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Backend server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
