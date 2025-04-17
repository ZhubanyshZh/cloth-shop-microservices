package controllers

import (
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/models"
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name"`
}

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input RegisterInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashedPassword, _ := utils.HashPassword(input.Password)

		user := models.User{
			Email:        input.Email,
			PasswordHash: hashedPassword,
			Name:         input.Name,
			Role:         models.RoleUser,
		}

		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User creation failed"})
			return
		}

		access, refresh, _ := utils.GenerateTokens(user.Email)
		c.JSON(http.StatusCreated, gin.H{"access": access, "refresh": refresh})
	}
}
