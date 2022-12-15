package config

import (
	"os"
	student "relationship/services/belongs-to/models"
	club "relationship/services/many-to-many/models"
	company "relationship/services/one-to-many/models"
	card "relationship/services/one-to-one/models"

	// a "relationship/services/one-to-many/models"

	env "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	err := env.Load(".env")
	if err != nil {
		panic(err)
	}

	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")
	POSTGRES_PORT := os.Getenv("POSTGRES_PORT")
	POSTGRES_SSL := os.Getenv("POSTGRES_SSL")
	POSTGRES_TIMEZONE := os.Getenv("POSTGRES_TIMEZONE")

	dsn := "host=" + POSTGRES_HOST + " " + "user=" + POSTGRES_USER + " " + "password=" + " " + POSTGRES_PASSWORD + " " +
		"dbname=" + POSTGRES_DB + " " + "port=" + POSTGRES_PORT + " " + "sslmode=" + POSTGRES_SSL + " " +
		"TimeZone=" + POSTGRES_TIMEZONE
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	//belongs to relationship
	db.AutoMigrate(&student.User{})
	db.AutoMigrate(&student.Company{})

	//one to one relationship
	db.AutoMigrate(&card.CreditCard{})
	db.AutoMigrate(&card.Customer{})

	//one to many relationship
	db.AutoMigrate(&company.Book{})
	db.AutoMigrate(&company.Page{})

	//many to many relationship normal
	db.AutoMigrate(&club.Club{})
	db.AutoMigrate(&club.Members{})

	DB = db
}
