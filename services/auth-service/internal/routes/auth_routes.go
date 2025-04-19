package routes

import (
	"fmt"
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/controllers"
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/middlewares"
	"github.com/gin-gonic/gin"
	"os"
)

func SetupAuthRoutes(controller controllers.AuthController) *gin.Engine {
	apiVersion := os.Getenv("API_VERSION")
	baseURL := fmt.Sprintf("/api/%s/auth", apiVersion)

	r := gin.Default()

	authRoute := r.Group(baseURL)
	{
		authRoute.POST("/register", middlewares.AuthReqMiddleware(), controller.Register)
		authRoute.POST("/login", middlewares.AuthReqMiddleware(), controller.Login)
	}

	return r
}
