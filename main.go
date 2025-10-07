package main

import (
	"log"
	"os"

	"github.com/Gthamsrim1/forms-backend/db"
	"github.com/Gthamsrim1/forms-backend/handlers"
	"github.com/Gthamsrim1/forms-backend/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	if err := db.Connect(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	r := gin.Default()

	public := r.Group("/api")
	{
		public.POST("/register", handlers.Register)
		public.POST("/login", handlers.Login)
	}

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", handlers.Profile)
		protected.POST("/create", handlers.CreateForm)
	}

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8000"
	}

	if err := r.Run(addr); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
