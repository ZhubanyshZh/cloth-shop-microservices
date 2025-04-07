package models

import (
	"time"
)

type Image struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProductID uint      `gorm:"index" json:"product_id"`
	URL       string    `gorm:"type:text;not null" json:"url"`
	IsMain    bool      `gorm:"default:false" json:"is_main"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
