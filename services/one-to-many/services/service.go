package services

import (
	"relationship/config"
	model "relationship/services/one-to-many/models"
)

func CreateData() {
	ab := model.G{
		Name: "aaa",
	}
	config.DB.Create(&ab)
}
