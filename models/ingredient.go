package models

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model
	Name    string   `json:"name"`
	Recipes []Recipe `gorm:"many2many:recipes_ingredients;"`
}
