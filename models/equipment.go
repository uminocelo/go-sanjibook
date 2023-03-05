package models

import "gorm.io/gorm"

type Equipment struct {
	gorm.Model
	Name    string   `json:"name"`
	Size    string   `json:"size"`
	Recipes []Recipe `gorm:"many2many:recipe_equipments"`
}
