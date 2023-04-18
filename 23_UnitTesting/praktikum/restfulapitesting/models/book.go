package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Id       string `json:"id" form:"id"`
	Judul    string `json:"judul" form:"judul"`
	Penulis  string `json:"penulis" form:"penulis"`
	Penerbit string `json:"penerbit" form:"penerbit"`
}
