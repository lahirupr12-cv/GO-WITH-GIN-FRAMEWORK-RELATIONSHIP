package manytomany

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company
}

type Company struct {
	gorm.Model
	Name string
}
