package repositories

import (
	"github.com/ZhubanyshZh/go-project-service/internal/models"
	"gorm.io/gorm"
)

type ImageRepository struct {
	DB *gorm.DB
}

func NewImageRepository(db *gorm.DB) *ImageRepository {
	return &ImageRepository{DB: db}
}

func (repo *ImageRepository) Create(image *models.Image) error {
	return repo.DB.Create(image).Error
}
