package main

import (
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/config"
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/config/db"
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/controllers"
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/repositories"
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/routes"
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/services"
	"log"
	"os"
)

func main() {
	config.LoadEnv()
	db.InitDB()

	authRepository := *repositories.NewImageRepository(db.DB)
	authService := *services.NewAuthService(authRepository)
	authController := *controllers.NewAuthController(authService)
	r := routes.SetupAuthRoutes(authController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("🚀 Server started on: ", port)
	log.Fatal(r.Run(":" + port))
}
