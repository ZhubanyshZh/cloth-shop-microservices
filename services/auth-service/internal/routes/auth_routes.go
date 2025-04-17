package routes

import (
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthRoutes(r *gin.Engine, db *gorm.DB) {
	auth := r.Group("/auth")
	auth.POST("/register", controllers.Register(db))
}
