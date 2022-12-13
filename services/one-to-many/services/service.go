package services

import (
	"relationship/config"
	model "relationship/services/one-to-many/models"
)

func CreateData() {
	c1 := model.Page{
		Number: "1",
	}
	c2 := model.Page{
		Number: "2",
	}
	c3 := model.Page{
		Number: "3",
	}
	u1 := model.Book{
		Name: "1",
		Page: []model.Page{c1, c2, c3},
	}
	config.DB.Create(&u1)
}
