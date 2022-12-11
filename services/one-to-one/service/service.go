package manytomany

import (
	"relationship/config"
	models "relationship/services/one-to-one/models"
)

func CreateData() {
	c1 := models.Company{
		Name: "company 1",
	}
	c2 := models.Company{
		Name: "company 2",
	}
	u1 := models.User{
		Name:    "lahiru",
		Company: c1,
	}
	u2 := models.User{
		Name:    "prasad",
		Company: c2,
	}

	config.DB.Create(&u1)
	config.DB.Create(&u2)
}
