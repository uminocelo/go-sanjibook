package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Name         string `gorm:"not null" json:"name"`
	Title        string `gorm:"not null" json:"title"`
	Description  string `gorm:"not null" json:"description"`
	Image        string `json:"image"`
	Servings     int    `gorm:"not null" json:"servings"`
	TotalTime    string `gorm:"not null" json:"total_time"`
	Rate         int    `json:"rate"`
	Tips         []Tip
	Instructions []Instruction
	Equipments   []Equipment  `gorm:"many2many:recipe_equipments;"`
	Ingredients  []Ingredient `gorm:"many2many:recipes_ingredients;"`
}
