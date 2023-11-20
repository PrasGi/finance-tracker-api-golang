package models

import "gorm.io/gorm"

type PersonalAccessToken struct {
	gorm.Model
	UserId uint
	Token  string `gorm:"unique"`
}
