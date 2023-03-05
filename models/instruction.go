package models

import "gorm.io/gorm"

type Instruction struct {
	gorm.Model
	Step        int    `json:"step"`
	Description string `json:"description"`
}
