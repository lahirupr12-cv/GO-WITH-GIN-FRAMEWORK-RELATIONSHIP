package models

import "gorm.io/gorm"

// User has and belongs to many languages, use `user_languages` as join table
type Class struct {
	gorm.Model
	Name    string
	Student []Student `gorm:"many2many:student_classes;"`
}

type Student struct {
	gorm.Model
	Name  string
	Class []Class `gorm:"many2many:student_classes;"`
}
