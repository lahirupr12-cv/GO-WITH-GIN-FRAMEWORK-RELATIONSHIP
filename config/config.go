package config

import (
	"os"
	student "relationship/services/one-to-one"

	env "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connection() {
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

	db.AutoMigrate(&student.Student{})
	DB = db
}
