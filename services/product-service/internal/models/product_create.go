package models

type ProductCreate struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `validate:"required"`
	Description string  `validate:"required"`
	Price       float64 `validate:"required,gt=0"`
}
