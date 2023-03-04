package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Name  string `gorm:"" json:"name"`
	Title string `gorm:"" json:"title"`
}
