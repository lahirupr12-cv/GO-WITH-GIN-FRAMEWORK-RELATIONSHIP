package services

import (
	"relationship/config"
	model "relationship/services/many-to-many/models"
)

func CreateData() {
	m1 := model.Members{
		Name: "lahiru",
	}
	m2 := model.Members{
		Name: "prasad",
	}
	m3 := model.Members{
		Name: "bandara",
	}
	c1 := model.Club{
		Members: []model.Members{m1, m2, m3},
	}

	config.DB.Create(&c1)
}
