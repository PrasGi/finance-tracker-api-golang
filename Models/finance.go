package models

import "gorm.io/gorm"

type Finance struct {
	gorm.Model
	Title      string
	Amount     float64
	UserId     uint     `gorm:"user_id"`
	User       User     `gorm:"foreignKey:UserId"`
	CategoryId uint     `gorm:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryId"`
}
