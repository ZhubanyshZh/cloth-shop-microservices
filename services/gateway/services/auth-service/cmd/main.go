package main

import (
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/config"
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/config/db"
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/routes"
	"log"
	"os"
)

func main() {
	config.LoadEnv()
	db.InitDB()
	r := routes.SetupAuthRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("🚀 Server started on: ", port)
	log.Fatal(r.Run(":" + port))
}
