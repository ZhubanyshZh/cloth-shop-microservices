package repositories

import (
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewImageRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(email, password string) (*models.User, error) {
	user := &models.User{Email: email, PasswordHash: password, Role: models.RoleUser}
	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) SaveRefreshToken(userID uuid.UUID, token string) error {
	rt := models.RefreshToken{
		UserID:    userID,
		Token:     token,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}
	return r.DB.Create(&rt).Error
}

func (r *UserRepository) DeleteRefreshToken(token string) error {
	return r.DB.Where("token = ?", token).Delete(&models.RefreshToken{}).Error
}
