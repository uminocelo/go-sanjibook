package models

import "gorm.io/gorm"

type Equipment struct {
	gorm.Model
	Name    string   `json:"name" gorm:"text"`
	Size    string   `json:"size" gorm:"text"`
	Recipes []Recipe `gorm:"many2many:recipe_equipments"`
}
