package models

import "gorm.io/gorm"

/*
	The only difference between hasOne and belongsTo is where the foreign key column is located. Let's say you have two entities: User and an Account.
	In short hasOne and belongsTo are inverses of one another - if one record belongTo the other, the other hasOne of the first.
*/

/*
	For example, if your application includes users and credit cards, and each user can only have one credit card.
*/
type Customer struct {
	gorm.Model
	Name       string
	CreditCard CreditCard `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;polymorphic:Owner"`
}

type CreditCard struct {
	gorm.Model
	Number    string
	OwnerID   uint
	OwnerType string
}
