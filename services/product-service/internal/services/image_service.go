package services

import (
	"github.com/ZhubanyshZh/go-project-service/internal/models"
	"gorm.io/gorm"
)

type ImageService struct {
	DB *gorm.DB
}

func NewImageService(db *gorm.DB) *ImageService {
	return &ImageService{DB: db}
}

func (s *ImageService) SaveImages(urls []string, productId uint) error {
	for _, url := range urls {
		image := &models.Image{
			ProductID: productId,
			URL:       url,
		}
		if err := s.DB.Create(image); err != nil {
			return err.Error
		}
	}
	return nil
}

func (s *ImageService) Create(image *models.Image) error {
	return s.DB.Create(image).Error
}
