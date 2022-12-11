package manytomany

import (
	"relationship/config"
	models "relationship/services/one-to-one/models"
)

func CreateStudent() {
	s1 := models.Student{
		Name: "lahiru",
	}

	config.DB.Create(&s1)
}
