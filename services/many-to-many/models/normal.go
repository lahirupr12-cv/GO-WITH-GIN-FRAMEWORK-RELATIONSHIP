package models

import "gorm.io/gorm"

// User has and belongs to many languages, `user_languages` is the join table
type Club struct {
	gorm.Model
	Members []Members `gorm:"many2many:club_members;"`
}

type Members struct {
	gorm.Model
	Name string
}
