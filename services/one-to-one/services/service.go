package services

import (
	"relationship/config"
	"relationship/services/one-to-one/models"
)

func CreateData() {
	card := models.CreditCard{
		Number: "9078181",
	}

	cus := models.Customer{
		Name:       "Lahiru",
		CreditCard: card,
	}

	config.DB.Create(&cus)

}
