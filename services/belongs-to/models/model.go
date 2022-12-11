package manytomany

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company `gorm:"constraint:onUpdate:CASCADE,OneDelete:SET NULL"`
}

type Company struct {
	gorm.Model
	Name string
}
