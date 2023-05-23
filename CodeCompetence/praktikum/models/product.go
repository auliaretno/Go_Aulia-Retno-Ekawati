package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	CategoryID  int      `json:"category_id" form:"category_id"`
	Category    Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Merk        string   `json:"merk" form:"merk"`
	Description string   `json:"description" form:"description"`
	Stock       int      `json:"stock" form:"stock"`
	Price       int      `json:"price" form:"price"`
}
