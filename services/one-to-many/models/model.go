package models

import "gorm.io/gorm"

type G struct {
	gorm.Model
	Name string
}
