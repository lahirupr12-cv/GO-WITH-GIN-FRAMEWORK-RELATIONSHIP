package models

import "gorm.io/gorm"

/*
	The only difference between hasOne and belongsTo is where the foreign key column is located. Let's say you have two entities: User and an Account.
	In short hasOne and belongsTo are inverses of one another - if one record belongTo the other, the other hasOne of the first.
*/

type Customer struct {
	gorm.Model
	CreditCard CreditCard
}

type CreditCard struct {
	gorm.Model
	Number     string
	CustomerID uint
}
