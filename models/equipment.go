package models

import "gorm.io/gorm"

type Equipment struct {
	gorm.Model
	Name    string   `gorm:"" json:"name"`
	Size    string   `gorm:"" json:"size"`
	Recipes []Recipe `gorm:"many2many:recipe_equipments"`
}
