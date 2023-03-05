package models

import "gorm.io/gorm"

type Tip struct {
	gorm.Model
	Description string `json:"description"`
}
