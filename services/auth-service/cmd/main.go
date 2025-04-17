package main

import (
	"fmt"
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/config"
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/models"
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config.LoadEnv()

	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	db.Exec(`DO $$ BEGIN
	    CREATE TYPE role_enum AS ENUM ('User', 'Admin');
	EXCEPTION
	    WHEN duplicate_object THEN null;
	END $$;`)

	db.AutoMigrate(&models.User{})

	r := gin.Default()
	routes.SetupAuthRoutes(r, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err = r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
