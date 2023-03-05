package models

import "gorm.io/gorm"

type Instruction struct {
	gorm.Model
	Step        int    `gorm:"" json:"step"`
	Description string `gorm:"" json:"description"`
}
