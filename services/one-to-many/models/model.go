package models

import "gorm.io/gorm"

/*
	For example, if your application includes Book and Pages, and each book can have many pages.
*/

type Book struct {
	gorm.Model
	Name string
	Page []Page
	// Page []Page `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;polymorphic:Book"`
}

type Page struct {
	gorm.Model
	Number string
	BookID uint
}
