package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `gorm:"coloumn:name;unique;not null"`
	Description string `gorm:"coloumn:description;not null"`
}
