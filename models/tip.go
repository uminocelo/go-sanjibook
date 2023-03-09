package models

import "gorm.io/gorm"

type Tip struct {
	gorm.Model
	Description string `gorm:"type:text" json:"description"`
	RecipeID    uint
}
