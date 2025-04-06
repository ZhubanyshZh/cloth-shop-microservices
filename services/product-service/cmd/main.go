package main

import (
	"fmt"
	"github.com/ZhubanyshZh/go-project-service/internal/cache/product_cache"
	"log"
	"net/http"
	"os"

	"github.com/ZhubanyshZh/go-project-service/internal/cache"
	"github.com/ZhubanyshZh/go-project-service/internal/config/db"
	"github.com/ZhubanyshZh/go-project-service/internal/handlers"
	"github.com/ZhubanyshZh/go-project-service/internal/repositories"
	"github.com/ZhubanyshZh/go-project-service/internal/routes"
	"github.com/ZhubanyshZh/go-project-service/internal/services"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db.InitDB()
	cache.InitRedis()

	productCache := product_cache.NewProductCache()
	repo := repositories.NewProductRepository(db.DB)
	service := services.NewProductService(repo, productCache)
	handler := handlers.ProductHandler{Service: service}

	r := routes.RegisterRoutes(&handler)

	port := os.Getenv("PORT")
	fmt.Println("🚀 Сервер запущен на порту", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
