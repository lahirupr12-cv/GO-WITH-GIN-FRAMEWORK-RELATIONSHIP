package manytomany

import (
	"gorm.io/gorm"
)

/*
For example, if your application includes users and companies, and each user can be assigned to exactly one company, the following types represent that relationship
*/
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
