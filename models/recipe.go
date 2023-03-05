package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Name         string `gorm:"" json:"name"`
	Title        string `gorm:"" json:"title"`
	Description  string `gorm:"" json:"description"`
	Image        string `gorm:"" json:"image"`
	Servings     int    `gorm:"" json:"servings"`
	TotalTime    string `gorm:"" json:"total_time"`
	Rate         int    `gorm:"" json:"rate"`
	Tips         []Tip
	Instructions []Instruction
	Equipments   []Equipment  `gorm:"many2many:recipe_equipments"`
	Ingredients  []Ingredient `gorm:"many2many:recipes_ingredients;"`
}
