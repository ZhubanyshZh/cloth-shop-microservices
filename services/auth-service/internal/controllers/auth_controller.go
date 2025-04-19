package controllers

import (
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/dtos"
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	Service *services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{Service: &service}
}

func (c *AuthController) Register(ctx *gin.Context) {
	userRegisterReqAny, exists := ctx.Get("validatedUserReq")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "❌ Invalid register req",
		})
		return
	}
	userRegisterReq := userRegisterReqAny.(dtos.AuthReq)
	user, err := c.Service.Register(userRegisterReq.Email, userRegisterReq.Password)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"user": user})
}

func (c *AuthController) Login(ctx *gin.Context) {
	userLoginReqAny, exists := ctx.Get("validatedUserReq")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "❌ Invalid login req",
		})
		return
	}
	userLoginReq := userLoginReqAny.(dtos.AuthReq)
	access, refresh, err := c.Service.Login(userLoginReq.Email, userLoginReq.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  access,
		"refresh_token": refresh,
	})
}
