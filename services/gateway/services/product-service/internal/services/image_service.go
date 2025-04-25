package services

import (
	"github.com/ZhubanyshZh/go-project-service/internal/repositories"
)

type ImageService struct {
	Repo repositories.ImageRepository
}

func NewImageService(repo repositories.ImageRepository) *ImageService {
	return &ImageService{Repo: repo}
}
