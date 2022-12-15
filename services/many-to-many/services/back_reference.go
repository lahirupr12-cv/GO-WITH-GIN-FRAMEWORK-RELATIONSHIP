package services

import (
	"relationship/config"
	"relationship/services/many-to-many/models"
)

func CreateOneSide() {
	p1 := models.Class{
		Name: "cls 1",
	}
	p2 := models.Class{
		Name: "cls 2",
	}
	p3 := models.Class{
		Name: "cls 3",
	}

	s := models.Student{
		Name: "Lahiru",
		Class: []models.Class{
			p1, p2, p3,
		},
	}

	config.DB.Create(&s)

}

func CreateOtherSide() {
	s1 := models.Student{
		Name: "Lahiru",
	}
	s2 := models.Student{
		Name: "prasad",
	}
	s3 := models.Student{
		Name: "banara",
	}

	c := models.Class{
		Name: "Sipara",
		Student: []models.Student{
			s1, s2, s3,
		},
	}

	config.DB.Create(&c)
}
