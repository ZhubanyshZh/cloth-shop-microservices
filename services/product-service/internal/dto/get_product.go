package dto

import (
	"github.com/ZhubanyshZh/go-project-service/internal/models"
	"time"
)

type GetProduct struct {
	ID          uint           `json:"id"`
	Name        string         `json:"product_name"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	CreatedAt   time.Time      `json:"created_at"`
	Images      []models.Image `json:"images"`
}
