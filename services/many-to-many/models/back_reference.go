package models

import "gorm.io/gorm"

// User has and belongs to many languages, use `user_languages` as join table
type PrivateClass struct {
	gorm.Model
	Name      string
	Languages []*Students `gorm:"many2many:class_studentes;"`
}

type Students struct {
	gorm.Model
	Name  string
	Users []*PrivateClass `gorm:"many2many:class_studentes;"`
}
