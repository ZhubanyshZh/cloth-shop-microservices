package repositories

import (
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/config/db"
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/models"
)

func CreateUser(email, password string) (*models.User, error) {
	user := &models.User{Email: email, PasswordHash: password, Role: models.RoleUser}
	if err := db.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func SaveUser(user *models.User) error {
	if err := db.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}
