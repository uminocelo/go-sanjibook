package models

import "gorm.io/gorm"

type RecipesIngredients struct {
	gorm.Model
	RecipeID     int    `gorm:"primaryKey" json:"recipe_id"`
	IngredientID int    `gorm:"primaryKey" json:"ingredient_id"`
	Measurement  string `gorm:"" json:"measurement"`
}
